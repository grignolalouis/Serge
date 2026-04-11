#!/usr/bin/env bash
# ============================================================================
# SERGE Evaluation Pipeline
# Runs 15 scenarios × N runs each, judges each run, then aggregates results.
#
# Usage:
#   ./eval/run.sh                     # full pipeline (45 agent + 45 judge + aggregation)
#   ./eval/run.sh --scenario uc1-01   # single scenario (3 runs)
#   ./eval/run.sh --dry-run           # show what would run without executing
#   ./eval/run.sh --judge-only        # skip agent runs, re-judge existing results
#   ./eval/run.sh --aggregate-only    # skip agent+judge, re-aggregate
# ============================================================================

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
EVAL_DIR="$SCRIPT_DIR"
SCENARIOS_DIR="$EVAL_DIR/scenarios"
RESULTS_DIR="$EVAL_DIR/results"
REPORT_DIR="$EVAL_DIR/report"

CONFIG="$EVAL_DIR/config.json"
JUDGE_PROMPT="$EVAL_DIR/judge_prompt.md"
JUDGE_SCHEMA="$EVAL_DIR/judge_schema.json"
REPORT_SCHEMA="$EVAL_DIR/report_schema.json"
AGGREGATION_PROMPT="$EVAL_DIR/aggregation_prompt.md"
MCP_CONFIG="$PROJECT_DIR/.mcp.json"
CLAUDE_MD="$PROJECT_DIR/CLAUDE.md"

RUNS_PER_SCENARIO=$(jq -r '.runs_per_scenario' "$CONFIG")
MODEL=$(jq -r '.model' "$CONFIG")
JUDGE_MODEL=$(jq -r '.judge_model' "$CONFIG")
AGGREGATION_MODEL=$(jq -r '.aggregation_model' "$CONFIG")
MAX_BUDGET=$(jq -r '.max_budget_per_run_usd' "$CONFIG")

FILTER_SCENARIO=""
DRY_RUN=false
JUDGE_ONLY=false
AGGREGATE_ONLY=false

while [[ $# -gt 0 ]]; do
  case "$1" in
    --scenario)    FILTER_SCENARIO="$2"; shift 2 ;;
    --dry-run)     DRY_RUN=true; shift ;;
    --judge-only)  JUDGE_ONLY=true; shift ;;
    --aggregate-only) AGGREGATE_ONLY=true; shift ;;
    *) echo "Unknown option: $1"; exit 1 ;;
  esac
done

timestamp() { date +"%Y-%m-%d %H:%M:%S"; }

log() { echo "[$(timestamp)] $*"; }

ensure_dirs() {
  mkdir -p "$RESULTS_DIR" "$REPORT_DIR"
}

# Step 1: Agent Run 
# Sends the scenario question to Claude with MCP access, captures full JSON output.
run_agent() {
  local scenario_file="$1"
  local output_file="$2"

  local question
  question=$(jq -r '.question' "$scenario_file")
  local scenario_id
  scenario_id=$(jq -r '.id' "$scenario_file")

  local system_context
  system_context=$(cat "$CLAUDE_MD")

  log "  AGENT [$scenario_id] → $output_file"

  if $DRY_RUN; then
    echo "{\"dry_run\": true, \"scenario\": \"$scenario_id\", \"question\": \"$question\"}" > "$output_file"
    return
  fi

  # Run claude in print mode with MCP server, capture JSON output
  local start_ms
  start_ms=$(python3 -c "import time; print(int(time.time()*1000))")

  claude -p "$question" \
    --model "$MODEL" \
    --mcp-config "$MCP_CONFIG" \
    --strict-mcp-config \
    --allowedTools "mcp__serge-supply-chain__*" \
    --append-system-prompt "$system_context" \
    --output-format json \
    --max-budget-usd "$MAX_BUDGET" \
    --no-session-persistence \
    --disable-slash-commands \
    > "$output_file.raw" 2>/dev/null || true

  local end_ms
  end_ms=$(python3 -c "import time; print(int(time.time()*1000))")
  local duration_ms=$(( end_ms - start_ms ))

  # Extract result and metadata from claude JSON output, wrap with our metadata
  local result
  result=$(jq -r '.result // empty' "$output_file.raw" 2>/dev/null || echo "")
  local cost
  cost=$(jq -r '.total_cost_usd // 0' "$output_file.raw" 2>/dev/null || echo "0")
  local num_turns
  num_turns=$(jq -r '.num_turns // 0' "$output_file.raw" 2>/dev/null || echo "0")
  local session_id
  session_id=$(jq -r '.session_id // ""' "$output_file.raw" 2>/dev/null || echo "")

  # Build our enriched result
  jq -n \
    --arg scenario_id "$scenario_id" \
    --arg question "$question" \
    --arg result "$result" \
    --argjson duration_ms "$duration_ms" \
    --argjson cost_usd "${cost:-0}" \
    --argjson num_turns "${num_turns:-0}" \
    --arg session_id "$session_id" \
    '{
      scenario_id: $scenario_id,
      question: $question,
      agent_answer: $result,
      duration_ms: $duration_ms,
      cost_usd: $cost_usd,
      num_turns: $num_turns,
      session_id: $session_id
    }' > "$output_file"

  rm -f "$output_file.raw"
  log "  AGENT [$scenario_id] done (${duration_ms}ms, \$${cost})"
}

# Step 2: Judge Run
# Sends the agent answer + ground truth to a judge LLM, gets structured scores.
run_judge() {
  local scenario_file="$1"
  local agent_result_file="$2"
  local judge_output_file="$3"

  local scenario_id
  scenario_id=$(jq -r '.id' "$scenario_file")
  local agent_answer
  agent_answer=$(jq -r '.agent_answer' "$agent_result_file")
  local ground_truth
  ground_truth=$(jq -c '.ground_truth' "$scenario_file")
  local expected_tools
  expected_tools=$(jq -c '.expected_tools' "$scenario_file")
  local title
  title=$(jq -r '.title' "$scenario_file")
  local difficulty
  difficulty=$(jq -r '.difficulty' "$scenario_file")

  local judge_system
  judge_system=$(cat "$JUDGE_PROMPT")

  log "  JUDGE [$scenario_id] → $judge_output_file"

  if $DRY_RUN; then
    echo "{\"dry_run\": true, \"scenario\": \"$scenario_id\"}" > "$judge_output_file"
    return
  fi

  local judge_input="## Scenario: ${title} (${difficulty})

## AGENT ANSWER:
${agent_answer}

## GROUND TRUTH:
${ground_truth}

## EXPECTED TOOLS:
${expected_tools}

Score this answer according to the rubric."

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
    > "$judge_output_file.raw" 2>/dev/null || true

  # Extract structured output: prefer .structured_output, fallback to parsing .result as JSON
  local extracted
  extracted=$(jq -e '.structured_output' "$judge_output_file.raw" 2>/dev/null) && {
    echo "$extracted" > "$judge_output_file"
  } || {
    # .result might be a JSON string or contain JSON
    local raw_result
    raw_result=$(jq -r '.result // ""' "$judge_output_file.raw" 2>/dev/null || echo "")
    # Try to parse as JSON directly
    if echo "$raw_result" | jq -e '.factual_accuracy' > /dev/null 2>&1; then
      echo "$raw_result" | jq '.' > "$judge_output_file"
    else
      # Extract JSON from markdown/text (find first { ... } block)
      local json_block
      json_block=$(echo "$raw_result" | python3 -c "
import sys, json, re
text = sys.stdin.read()
# Find JSON objects in the text
matches = re.findall(r'\{[^{}]*(?:\{[^{}]*\}[^{}]*)*\}', text)
for m in matches:
    try:
        obj = json.loads(m)
        if 'factual_accuracy' in obj:
            print(json.dumps(obj))
            sys.exit(0)
    except: pass
# Fallback: empty judge result
print(json.dumps({'factual_accuracy':0,'completeness':0,'reasoning':0,'source_attribution':0,'structure':0,'weighted_score':0,'pass':False,'justification':'Failed to parse judge output','facts_found':[],'facts_missing':[],'hallucinations':[]}))
" 2>/dev/null)
      echo "$json_block" > "$judge_output_file"
    fi
  }
  rm -f "$judge_output_file.raw"

  local score
  score=$(jq -r '.weighted_score // "?"' "$judge_output_file")
  local pass
  pass=$(jq -r '.pass // "?"' "$judge_output_file")
  log "  JUDGE [$scenario_id] done (score: $score, pass: $pass)"
}

# Step 3: Aggregation
run_aggregation() {
  log "AGGREGATION: collecting all results..."

  if $DRY_RUN; then
    echo "{\"dry_run\": true}" > "$REPORT_DIR/evaluation_report.json"
    log "AGGREGATION: dry run complete."
    return
  fi

  log "AGGREGATION: computing..."
  local report_file="$REPORT_DIR/evaluation_report.json"
  python3 "$EVAL_DIR/aggregate.py" "$SCENARIOS_DIR" "$RESULTS_DIR" "$RUNS_PER_SCENARIO" > "$report_file"

  log "AGGREGATION: report written to $REPORT_DIR/evaluation_report.json"
}

# Print Summary
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

  echo ""
  echo "  Total runs:        $total_runs"
  echo "  Avg weighted score: $avg_score / 5.0"
  echo "  Pass rate:         $pass_rate"
  echo "  Total cost:        \$$total_cost"
  echo ""

  echo "  By Use Case:"
  jq -r '.by_use_case[] | "    \(.use_case): score=\(.avg_score) pass=\(.pass_rate)"' "$report" 2>/dev/null || true

  echo ""
  echo "  By Difficulty:"
  jq -r '.by_difficulty[] | "    \(.difficulty): score=\(.avg_score) pass=\(.pass_rate)"' "$report" 2>/dev/null || true

  echo ""
}

# Main
main() {
  ensure_dirs

  log "SERGE Evaluation Pipeline"
  log "Config: $RUNS_PER_SCENARIO runs/scenario, model=$MODEL, judge=$JUDGE_MODEL"

  if $AGGREGATE_ONLY; then
    run_aggregation
    print_summary
    return
  fi

  # Collect scenarios to run
  local scenarios=()
  if [ -n "$FILTER_SCENARIO" ]; then
    local f="$SCENARIOS_DIR/${FILTER_SCENARIO}.json"
    if [ ! -f "$f" ]; then
      log "ERROR: scenario file not found: $f"
      exit 1
    fi
    scenarios+=("$f")
  else
    for f in "$SCENARIOS_DIR"/*.json; do
      scenarios+=("$f")
    done
  fi

  local total_scenarios=${#scenarios[@]}
  local total_runs=$((total_scenarios * RUNS_PER_SCENARIO))
  log "Running $total_runs evaluations ($total_scenarios scenarios × $RUNS_PER_SCENARIO runs)"

  local scenario_num=0
  for scenario_file in "${scenarios[@]}"; do
    scenario_num=$((scenario_num + 1))
    local scenario_id
    scenario_id=$(jq -r '.id' "$scenario_file")
    local title
    title=$(jq -r '.title' "$scenario_file")

    log "━━━ [$scenario_num/$total_scenarios] $scenario_id: $title ━━━"

    local result_dir="$RESULTS_DIR/$scenario_id"
    mkdir -p "$result_dir"

    for run in $(seq 1 "$RUNS_PER_SCENARIO"); do
      log "  Run $run/$RUNS_PER_SCENARIO"

      local agent_file="$result_dir/run_${run}.json"
      local judge_file="$result_dir/run_${run}_judge.json"

      # Step 1: Agent run (skip if --judge-only)
      if ! $JUDGE_ONLY; then
        run_agent "$scenario_file" "$agent_file"
      fi

      # Step 2: Judge run
      if [ -f "$agent_file" ]; then
        run_judge "$scenario_file" "$agent_file" "$judge_file"
      else
        log "  SKIP JUDGE: no agent result for run $run"
      fi
    done
  done

  # Step 3: Aggregation
  run_aggregation
  print_summary

  log "Pipeline complete."
}

main "$@"
