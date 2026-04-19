#!/usr/bin/env bash
# ============================================================================
# SERGE Evaluation Pipeline — resumable, rate-limit aware.
#
# The pipeline is idempotent: re-running skips any scenario/run whose output
# file already exists and is valid. Transient failures (network, rate limit)
# are retried with exponential backoff. Persistent failures are logged and
# skipped so the pipeline can finish and be resumed later.
#
# Usage:
#   ./eval/run.sh                   # full pipeline (resume if interrupted)
#   ./eval/run.sh --scenario L4-01  # single scenario
#   ./eval/run.sh --dry-run         # show what would run
#   ./eval/run.sh --judge-only      # re-judge existing agent results
#   ./eval/run.sh --aggregate-only  # re-aggregate only
#   ./eval/run.sh --force-retry     # re-run even if outputs already exist
#   ./eval/run.sh --status          # print how many runs are done / pending
#   ./eval/run.sh --label run2      # write to results_run2/ and report_run2/
# ============================================================================

set -uo pipefail  # note: removed -e to handle individual call failures

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
EVAL_DIR="$SCRIPT_DIR"
SCENARIOS_DIR="$EVAL_DIR/scenarios"

CONFIG="$EVAL_DIR/config.json"
JUDGE_PROMPT="$EVAL_DIR/judge_prompt.md"
JUDGE_SCHEMA="$EVAL_DIR/judge_schema.json"
MCP_CONFIG="$PROJECT_DIR/.mcp.json"
CLAUDE_MD="$PROJECT_DIR/CLAUDE.md"

RUNS_PER_SCENARIO=$(jq -r '.runs_per_scenario' "$CONFIG")
MODEL=$(jq -r '.model' "$CONFIG")
JUDGE_MODEL=$(jq -r '.judge_model' "$CONFIG")
MAX_BUDGET=$(jq -r '.max_budget_per_run_usd' "$CONFIG")

# Resilience config
MAX_RETRIES=3
BACKOFF_BASE=15      # seconds, exponential: 15, 45, 135
QUOTA_EXIT_CODE=2

FILTER_SCENARIO=""
DRY_RUN=false
JUDGE_ONLY=false
AGGREGATE_ONLY=false
FORCE_RETRY=false
STATUS_ONLY=false
LABEL=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --scenario)       FILTER_SCENARIO="$2"; shift 2 ;;
    --dry-run)        DRY_RUN=true; shift ;;
    --judge-only)     JUDGE_ONLY=true; shift ;;
    --aggregate-only) AGGREGATE_ONLY=true; shift ;;
    --force-retry)    FORCE_RETRY=true; shift ;;
    --status)         STATUS_ONLY=true; shift ;;
    --label)          LABEL="$2"; shift 2 ;;
    *) echo "Unknown option: $1"; exit 1 ;;
  esac
done

# Label-suffixed output dirs keep multiple runs isolated
if [ -n "$LABEL" ]; then
  RESULTS_DIR="$EVAL_DIR/results_${LABEL}"
  REPORT_DIR="$EVAL_DIR/report_${LABEL}"
else
  RESULTS_DIR="$EVAL_DIR/results"
  REPORT_DIR="$EVAL_DIR/report"
fi
FAILURE_LOG="$RESULTS_DIR/_failures.log"
QUOTA_MARKER="$RESULTS_DIR/_quota_hit"

timestamp() { date +"%Y-%m-%d %H:%M:%S"; }
log() { echo "[$(timestamp)] $*"; }
log_err() { echo "[$(timestamp)] ERROR: $*" >&2; echo "[$(timestamp)] $*" >> "$FAILURE_LOG"; }

ensure_dirs() { mkdir -p "$RESULTS_DIR" "$REPORT_DIR"; }

# Signal handling — clean exit on Ctrl+C
trap 'log ""; log "Interrupted. Progress is preserved — re-run the script to resume."; exit 130' INT TERM

# ── Validation helpers ──────────────────────────────────────────────────
is_valid_agent_result() {
  local f="$1"
  [ -f "$f" ] || return 1
  local answer
  answer=$(jq -r '.agent_answer // ""' "$f" 2>/dev/null) || return 1
  [ -n "$answer" ] && [ "$answer" != "null" ]
}

is_valid_judge_result() {
  local f="$1"
  [ -f "$f" ] || return 1
  local score
  score=$(jq -r '.weighted_score // 0' "$f" 2>/dev/null) || return 1
  awk -v s="$score" 'BEGIN{exit !(s+0 > 0)}'
}

# Classify a claude CLI failure.
#
# Priority: explicit error patterns in stderr > explicit error events in stream.
# Informational events (rate_limit_event with status=allowed, rateLimitType metadata)
# are NOT treated as errors — they just describe the current subscription state.
#
# Returns: quota | network | auth | budget | parse | timeout | unknown
classify_error() {
  local stream_file="$1"
  local err_file="${2:-}"

  local stderr_tail=""
  local stream_tail=""
  [ -f "$err_file" ]    && stderr_tail="$(tail -c 8192 "$err_file" 2>/dev/null)"
  [ -f "$stream_file" ] && stream_tail="$(tail -c 8192 "$stream_file" 2>/dev/null)"

  # 1) stderr — when present, it is authoritative
  if [ -n "$stderr_tail" ]; then
    if echo "$stderr_tail" | grep -qiE "(rate[ _-]?limit|quota|usage[ _-]?limit|too many request|subscription (limit|reached)).*(exceed|reach|hit|block)"; then
      echo "quota"; return
    fi
    if echo "$stderr_tail" | grep -qiE "^Error.*(401|403|unauthor|authentication|invalid api key|not logged in)"; then
      echo "auth"; return
    fi
    if echo "$stderr_tail" | grep -qiE "(budget|cost).{0,40}(exceed|limit reached)"; then
      echo "budget"; return
    fi
    if echo "$stderr_tail" | grep -qiE "EAI_AGAIN|ENOTFOUND|ECONNRESET|ECONNREFUSED|socket hang up|getaddrinfo"; then
      echo "network"; return
    fi
    if echo "$stderr_tail" | grep -qiE "ETIMEDOUT|timed out|request timeout"; then
      echo "timeout"; return
    fi
  fi

  # 2) stream — look for explicit error events, not informational ones
  if [ -n "$stream_tail" ]; then
    # Actual quota block: the API returns an error event or a rate_limit_event with status=blocked
    if echo "$stream_tail" | grep -qE '"(is_error|error)":\s*true' \
       && echo "$stream_tail" | grep -qiE "rate.?limit|quota|too many request|overage"; then
      echo "quota"; return
    fi
    if echo "$stream_tail" | grep -qE '"status"\s*:\s*"(blocked|exceeded|denied)"'; then
      echo "quota"; return
    fi
    # Budget exhaustion (max-budget-usd) — Claude CLI reports it in the result event
    if echo "$stream_tail" | grep -qE '"subtype"\s*:\s*"(error_max_turns|budget_exceeded|max_budget_reached)"'; then
      echo "budget"; return
    fi
    if echo "$stream_tail" | grep -qE '"stop_reason"\s*:\s*"max_turns"'; then
      echo "budget"; return
    fi
  fi

  echo "unknown"
}

# Print a compact diagnostic of the last failure: exit code, stderr tail, stream tail.
dump_failure_context() {
  local scenario_id="$1"
  local kind="$2"              # AGENT | JUDGE
  local exit_code="$3"
  local stream_file="${4:-}"
  local err_file="${5:-}"

  {
    echo "---- $kind [$scenario_id] failure detail ($(timestamp)) ----"
    echo "exit code: $exit_code"
    if [ -f "$err_file" ] && [ -s "$err_file" ]; then
      echo "stderr (last 15 lines):"
      tail -n 15 "$err_file" | sed 's/^/  /'
    else
      echo "stderr: (empty)"
    fi
    if [ -f "$stream_file" ] && [ -s "$stream_file" ]; then
      echo "stream (last 15 lines):"
      tail -n 15 "$stream_file" | sed 's/^/  /'
    else
      echo "stream: (empty)"
    fi
    echo "-----------------------------------------"
  } | tee -a "$FAILURE_LOG"
}

check_quota_marker() {
  [ -f "$QUOTA_MARKER" ]
}

set_quota_marker() {
  date > "$QUOTA_MARKER"
}

clear_quota_marker() {
  rm -f "$QUOTA_MARKER"
}

# ── Agent run (with retry, resume, quota detection) ────────────────────
run_agent() {
  local scenario_file="$1"
  local output_file="$2"
  local scenario_id
  scenario_id=$(jq -r '.id' "$scenario_file")

  if ! $FORCE_RETRY && is_valid_agent_result "$output_file"; then
    log "  AGENT [$scenario_id] skip — already done"
    return 0
  fi

  if $DRY_RUN; then
    jq -n --arg id "$scenario_id" '{dry_run: true, scenario: $id}' > "$output_file"
    return 0
  fi

  local question
  question=$(jq -r '.question' "$scenario_file")
  local system_context
  system_context=$(cat "$CLAUDE_MD")

  local attempt=1
  local backoff=$BACKOFF_BASE
  while [ $attempt -le $MAX_RETRIES ]; do
    log "  AGENT [$scenario_id] attempt $attempt/$MAX_RETRIES"

    local start_ms
    start_ms=$(python3 -c "import time; print(int(time.time()*1000))")

    claude -p "$question" \
      --model "$MODEL" \
      --mcp-config "$MCP_CONFIG" \
      --strict-mcp-config \
      --allowedTools "mcp__serge-supply-chain__*" \
      --append-system-prompt "$system_context" \
      --output-format stream-json \
      --verbose \
      --max-budget-usd "$MAX_BUDGET" \
      --no-session-persistence \
      --disable-slash-commands \
      > "$output_file.stream" 2>"$output_file.err"
    local claude_exit=$?

    local end_ms
    end_ms=$(python3 -c "import time; print(int(time.time()*1000))")
    local duration_ms=$(( end_ms - start_ms ))

    # Parse stream into the final output file
    python3 - "$scenario_id" "$question" "$duration_ms" "$output_file.stream" "$output_file" <<'PYEOF'
import json, sys

scenario_id, question, duration_ms, stream_file, output_file = sys.argv[1:6]
duration_ms = int(duration_ms)

result_text = ""
cost = 0.0
num_turns = 0
session_id = ""
tools_used = []

try:
    with open(stream_file) as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            try:
                ev = json.loads(line)
            except json.JSONDecodeError:
                continue

            ev_type = ev.get("type", "")
            if ev_type == "assistant":
                content = ev.get("message", {}).get("content", [])
                if isinstance(content, list):
                    for block in content:
                        if isinstance(block, dict) and block.get("type") == "tool_use":
                            tools_used.append(block.get("name", ""))
            elif ev_type == "result":
                result_text = ev.get("result", "") or result_text
                cost = ev.get("total_cost_usd", cost) or cost
                num_turns = ev.get("num_turns", num_turns) or num_turns
                session_id = ev.get("session_id", session_id) or session_id
except FileNotFoundError:
    pass

system_prefixes = {"srm_": "SRM", "wms_": "WMS", "oms_": "OMS", "tms_": "TMS"}
systems_touched = set()
for name in tools_used:
    for prefix, sysname in system_prefixes.items():
        if prefix in name:
            systems_touched.add(sysname)
            break

out = {
    "scenario_id": scenario_id,
    "question": question,
    "agent_answer": result_text,
    "duration_ms": duration_ms,
    "cost_usd": float(cost),
    "num_turns": int(num_turns),
    "session_id": session_id,
    "tools_used": tools_used,
    "tool_calls_count": len(tools_used),
    "systems_touched": sorted(systems_touched),
}

with open(output_file, "w") as f:
    json.dump(out, f, indent=2)
PYEOF

    if is_valid_agent_result "$output_file"; then
      local tools_count
      tools_count=$(jq -r '.tool_calls_count' "$output_file")
      local systems
      systems=$(jq -r '.systems_touched | join(",")' "$output_file")
      log "  AGENT [$scenario_id] ok (${duration_ms}ms, tools=$tools_count, systems=${systems:-none})"
      rm -f "$output_file.stream" "$output_file.err"
      return 0
    fi

    # Failure — classify, log and decide
    local err_type
    err_type=$(classify_error "$output_file.stream" "$output_file.err")
    local stream_bytes err_bytes
    stream_bytes=$(wc -c < "$output_file.stream" 2>/dev/null | tr -d ' ')
    err_bytes=$(wc -c < "$output_file.err" 2>/dev/null | tr -d ' ')

    log_err "AGENT [$scenario_id] failed (attempt $attempt, exit=$claude_exit, class=$err_type, stream=${stream_bytes:-0}B, stderr=${err_bytes:-0}B)"
    dump_failure_context "$scenario_id" "AGENT" "$claude_exit" "$output_file.stream" "$output_file.err"

    if [ "$err_type" = "quota" ]; then
      set_quota_marker
      log_err "Quota/rate limit detected. Stopping — re-run when quota resets."
      return $QUOTA_EXIT_CODE
    fi

    if [ "$err_type" = "auth" ] || [ "$err_type" = "budget" ]; then
      log_err "Non-transient error ($err_type). Stopping — fix the configuration before retrying."
      return 1
    fi

    if [ $attempt -lt $MAX_RETRIES ]; then
      log "  AGENT [$scenario_id] backing off ${backoff}s before retry"
      sleep $backoff
      backoff=$(( backoff * 3 ))
    fi
    attempt=$(( attempt + 1 ))
  done

  log_err "AGENT [$scenario_id] gave up after $MAX_RETRIES attempts (diagnostics preserved at $output_file.stream / .err)"
  return 1
}

# ── Judge run (with retry, resume) ─────────────────────────────────────
run_judge() {
  local scenario_file="$1"
  local agent_result_file="$2"
  local judge_output_file="$3"
  local scenario_id
  scenario_id=$(jq -r '.id' "$scenario_file")

  if ! $FORCE_RETRY && is_valid_judge_result "$judge_output_file"; then
    log "  JUDGE [$scenario_id] skip — already done"
    return 0
  fi

  if ! is_valid_agent_result "$agent_result_file"; then
    log "  JUDGE [$scenario_id] skip — no valid agent result"
    return 1
  fi

  if $DRY_RUN; then
    jq -n --arg id "$scenario_id" '{dry_run: true, scenario: $id}' > "$judge_output_file"
    return 0
  fi

  local agent_answer ground_truth title level use_case
  agent_answer=$(jq -r '.agent_answer' "$agent_result_file")
  ground_truth=$(jq -c '.ground_truth' "$scenario_file")
  title=$(jq -r '.title' "$scenario_file")
  level=$(jq -r '.level' "$scenario_file")
  use_case=$(jq -r '.use_case' "$scenario_file")

  local judge_system
  judge_system=$(cat "$JUDGE_PROMPT")

  local judge_input="## Scenario: ${title} (${level}, ${use_case})

## AGENT ANSWER:
${agent_answer}

## GROUND TRUTH:
${ground_truth}

Score this answer according to the rubric. Remember: flexible on form, strict on facts."

  local attempt=1
  local backoff=$BACKOFF_BASE
  while [ $attempt -le $MAX_RETRIES ]; do
    log "  JUDGE [$scenario_id] attempt $attempt/$MAX_RETRIES"

    claude -p "$judge_input" \
      --model "$JUDGE_MODEL" \
      --system-prompt "$judge_system" \
      --output-format json \
      --json-schema "$(cat "$JUDGE_SCHEMA")" \
      --no-session-persistence \
      --max-budget-usd 0.10 \
      --disable-slash-commands \
      --allowedTools "" \
      --tools "" \
      > "$judge_output_file.raw" 2>"$judge_output_file.err"
    local claude_exit=$?

    # Extract structured output
    local extracted
    if extracted=$(jq -e '.structured_output' "$judge_output_file.raw" 2>/dev/null); then
      echo "$extracted" > "$judge_output_file"
    else
      local raw_result
      raw_result=$(jq -r '.result // ""' "$judge_output_file.raw" 2>/dev/null || echo "")
      if echo "$raw_result" | jq -e '.factual_accuracy' > /dev/null 2>&1; then
        echo "$raw_result" | jq '.' > "$judge_output_file"
      else
        echo "$raw_result" | python3 -c "
import sys, json, re
text = sys.stdin.read()
matches = re.findall(r'\{[^{}]*(?:\{[^{}]*\}[^{}]*)*\}', text)
for m in matches:
    try:
        obj = json.loads(m)
        if 'factual_accuracy' in obj:
            print(json.dumps(obj)); sys.exit(0)
    except: pass
sys.exit(1)
" > "$judge_output_file" 2>/dev/null || rm -f "$judge_output_file"
      fi
    fi

    if is_valid_judge_result "$judge_output_file"; then
      local score pass
      score=$(jq -r '.weighted_score // "?"' "$judge_output_file")
      pass=$(jq -r '.pass // "?"' "$judge_output_file")
      log "  JUDGE [$scenario_id] ok (score: $score, pass: $pass)"
      rm -f "$judge_output_file.raw" "$judge_output_file.err"
      return 0
    fi

    # Failure — classify, log and decide
    local err_type
    err_type=$(classify_error "$judge_output_file.raw" "$judge_output_file.err")
    local raw_bytes err_bytes
    raw_bytes=$(wc -c < "$judge_output_file.raw" 2>/dev/null | tr -d ' ')
    err_bytes=$(wc -c < "$judge_output_file.err" 2>/dev/null | tr -d ' ')

    log_err "JUDGE [$scenario_id] failed (attempt $attempt, exit=$claude_exit, class=$err_type, raw=${raw_bytes:-0}B, stderr=${err_bytes:-0}B)"
    dump_failure_context "$scenario_id" "JUDGE" "$claude_exit" "$judge_output_file.raw" "$judge_output_file.err"

    if [ "$err_type" = "quota" ]; then
      set_quota_marker
      log_err "Quota/rate limit detected during judge. Stopping — re-run when quota resets."
      return $QUOTA_EXIT_CODE
    fi

    if [ "$err_type" = "auth" ] || [ "$err_type" = "budget" ]; then
      log_err "Non-transient error ($err_type). Stopping — fix the configuration before retrying."
      return 1
    fi

    if [ $attempt -lt $MAX_RETRIES ]; then
      log "  JUDGE [$scenario_id] backing off ${backoff}s before retry"
      sleep $backoff
      backoff=$(( backoff * 3 ))
    fi
    attempt=$(( attempt + 1 ))
  done

  log_err "JUDGE [$scenario_id] gave up after $MAX_RETRIES attempts (diagnostics preserved at $judge_output_file.raw / .err)"
  return 1
}

# ── Status report ──────────────────────────────────────────────────────
print_status() {
  local total_agent=0 done_agent=0
  local total_judge=0 done_judge=0
  for scenario_file in "$SCENARIOS_DIR"/*.json; do
    local sid
    sid=$(jq -r '.id' "$scenario_file")
    for run in $(seq 1 "$RUNS_PER_SCENARIO"); do
      total_agent=$(( total_agent + 1 ))
      total_judge=$(( total_judge + 1 ))
      if is_valid_agent_result "$RESULTS_DIR/$sid/run_${run}.json"; then
        done_agent=$(( done_agent + 1 ))
      fi
      if is_valid_judge_result "$RESULTS_DIR/$sid/run_${run}_judge.json"; then
        done_judge=$(( done_judge + 1 ))
      fi
    done
  done
  echo ""
  echo "  SERGE Evaluation Status"
  echo "  Agent runs: $done_agent / $total_agent"
  echo "  Judge runs: $done_judge / $total_judge"
  if check_quota_marker; then
    echo "  ⚠ Quota marker present: $QUOTA_MARKER"
  fi
  echo ""
}

# ── Aggregation ────────────────────────────────────────────────────────
run_aggregation() {
  log "AGGREGATION: computing..."
  if $DRY_RUN; then
    echo "{\"dry_run\": true}" > "$REPORT_DIR/evaluation_report.json"
    return
  fi
  python3 "$EVAL_DIR/aggregate.py" "$SCENARIOS_DIR" "$RESULTS_DIR" "$RUNS_PER_SCENARIO" > "$REPORT_DIR/evaluation_report.json"
  log "AGGREGATION: report → $REPORT_DIR/evaluation_report.json"
}

print_summary() {
  local report="$REPORT_DIR/evaluation_report.json"
  [ -f "$report" ] || return
  echo ""
  echo "  SERGE EVALUATION REPORT"
  local avg_score pass_rate total_runs total_cost
  avg_score=$(jq -r '.summary.avg_weighted_score // "N/A"' "$report")
  pass_rate=$(jq -r '.summary.pass_rate // "N/A"' "$report")
  total_runs=$(jq -r '.summary.total_runs // "N/A"' "$report")
  total_cost=$(jq -r '.metrics.total_cost_usd // "N/A"' "$report")
  echo "  Total runs:         $total_runs"
  echo "  Avg weighted score: $avg_score / 5.0"
  echo "  Pass rate:          $pass_rate"
  echo "  Total cost:         \$$total_cost"
  echo ""
  echo "  By Level:"
  jq -r '.by_level[] | "    \(.level): score=\(.avg_score) pass=\(.pass_rate) avg_tools=\(.avg_tool_calls)"' "$report" 2>/dev/null || true
  echo ""
}

# ── Main ───────────────────────────────────────────────────────────────
main() {
  ensure_dirs

  if $STATUS_ONLY; then
    print_status
    return 0
  fi

  log "SERGE Evaluation Pipeline (resumable)"
  log "Config: $RUNS_PER_SCENARIO runs/scenario, model=$MODEL, judge=$JUDGE_MODEL"

  if check_quota_marker && ! $AGGREGATE_ONLY && ! $FORCE_RETRY; then
    log "⚠ Previous run hit a quota limit. Delete $QUOTA_MARKER to proceed, or use --force-retry."
    log "  Current progress:"
    print_status
    exit $QUOTA_EXIT_CODE
  fi

  if $AGGREGATE_ONLY; then
    run_aggregation
    print_summary
    return 0
  fi

  local scenarios=()
  if [ -n "$FILTER_SCENARIO" ]; then
    local f="$SCENARIOS_DIR/${FILTER_SCENARIO}.json"
    [ -f "$f" ] || { log_err "scenario not found: $f"; exit 1; }
    scenarios+=("$f")
  else
    for f in "$SCENARIOS_DIR"/*.json; do
      scenarios+=("$f")
    done
  fi

  local total_scenarios=${#scenarios[@]}
  local total_runs=$((total_scenarios * RUNS_PER_SCENARIO))
  log "Running $total_runs evaluations ($total_scenarios scenarios x $RUNS_PER_SCENARIO runs)"

  # Counters
  local agents_done=0 agents_skipped=0 agents_failed=0
  local judges_done=0 judges_skipped=0 judges_failed=0

  local scenario_num=0
  for scenario_file in "${scenarios[@]}"; do
    scenario_num=$((scenario_num + 1))
    local scenario_id title
    scenario_id=$(jq -r '.id' "$scenario_file")
    title=$(jq -r '.title' "$scenario_file")

    log "━━━ [$scenario_num/$total_scenarios] $scenario_id: $title ━━━"

    local result_dir="$RESULTS_DIR/$scenario_id"
    mkdir -p "$result_dir"

    for run in $(seq 1 "$RUNS_PER_SCENARIO"); do
      local agent_file="$result_dir/run_${run}.json"
      local judge_file="$result_dir/run_${run}_judge.json"

      # Agent
      if ! $JUDGE_ONLY; then
        local agent_was_done=false
        if is_valid_agent_result "$agent_file" && ! $FORCE_RETRY; then
          agent_was_done=true
        fi

        run_agent "$scenario_file" "$agent_file"
        local rc=$?
        if [ $rc -eq $QUOTA_EXIT_CODE ]; then
          log ""
          log "Stopping pipeline: quota reached. Re-run the script to resume."
          print_status
          exit $QUOTA_EXIT_CODE
        elif [ $rc -eq 0 ]; then
          if $agent_was_done; then
            agents_skipped=$((agents_skipped + 1))
          else
            agents_done=$((agents_done + 1))
          fi
        else
          agents_failed=$((agents_failed + 1))
        fi
      fi

      # Judge
      if is_valid_agent_result "$agent_file"; then
        local judge_was_done=false
        if is_valid_judge_result "$judge_file" && ! $FORCE_RETRY; then
          judge_was_done=true
        fi

        run_judge "$scenario_file" "$agent_file" "$judge_file"
        local rc=$?
        if [ $rc -eq $QUOTA_EXIT_CODE ]; then
          log ""
          log "Stopping pipeline: quota reached. Re-run the script to resume."
          print_status
          exit $QUOTA_EXIT_CODE
        elif [ $rc -eq 0 ]; then
          if $judge_was_done; then
            judges_skipped=$((judges_skipped + 1))
          else
            judges_done=$((judges_done + 1))
          fi
        else
          judges_failed=$((judges_failed + 1))
        fi
      else
        log "  JUDGE [$scenario_id] skip — no valid agent result for run $run"
        judges_failed=$((judges_failed + 1))
      fi
    done
  done

  clear_quota_marker

  log ""
  log "Pipeline pass complete."
  log "  Agents:  $agents_done new, $agents_skipped skipped, $agents_failed failed"
  log "  Judges:  $judges_done new, $judges_skipped skipped, $judges_failed failed"

  run_aggregation
  print_summary

  if [ $agents_failed -gt 0 ] || [ $judges_failed -gt 0 ]; then
    log ""
    log "Some items failed. See $FAILURE_LOG for details. Re-run to retry pending items."
  fi
}

main "$@"
