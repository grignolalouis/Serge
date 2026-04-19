# Pipeline Commands — SERGE Evaluation

Reference for running the evaluation pipeline (`eval/run.sh`).

Run commands from the `serge/` directory.

## Quick reference

| Command | Purpose |
|---------|---------|
| `./eval/run.sh` | Run the full pipeline (or resume if interrupted) |
| `./eval/run.sh --status` | Show progress without launching any run |
| `./eval/run.sh --scenario L4-01` | Run a single scenario (3 runs by default) |
| `./eval/run.sh --judge-only` | Re-judge existing agent results |
| `./eval/run.sh --aggregate-only` | Rebuild the report from existing results |
| `./eval/run.sh --force-retry` | Re-run items even if outputs already exist |
| `./eval/run.sh --dry-run` | Show what would run, no API calls |

## Typical workflow

### Starting from scratch
```bash
cd serge
./eval/run.sh
```
Runs all 27 scenarios × 3 runs = 81 agent runs + 81 judge runs, then aggregates the report.

### Resuming after interrupt or quota hit
```bash
./eval/run.sh
```
The pipeline is **idempotent** — any run with a valid output file is skipped. Progress counters at the end show what was new vs. skipped vs. failed.

### Checking progress without running
```bash
./eval/run.sh --status
```
Example output:
```
  SERGE Evaluation Status
  Agent runs: 45 / 81
  Judge runs: 42 / 81
```

### Running a single scenario
```bash
./eval/run.sh --scenario L4-01
```
Useful for debugging one scenario. Runs 3 iterations and judges each.

Scenario IDs follow `L[1-5]-[01-06]` format — see `eval/scenarios/`.

### Re-running only the judges
```bash
./eval/run.sh --judge-only
```
If you've updated `judge_prompt.md` or `judge_schema.json` and want to re-score existing agent answers without paying for new agent runs.

### Re-aggregating only
```bash
./eval/run.sh --aggregate-only
```
If you've updated `aggregate.py` logic and want to rebuild `eval/report/evaluation_report.json` from existing judge results.

### Forcing a re-run
```bash
./eval/run.sh --force-retry
./eval/run.sh --force-retry --scenario L4-01
```
Overrides the skip-if-valid logic. Use sparingly — costs extra quota.

## Handling quota limits

When the Claude subscription quota is hit, the pipeline:

1. Detects `rate limit`, `quota`, `429`, or similar patterns in Claude's output
2. Creates a marker file: `eval/results/_quota_hit`
3. Exits with code **2** and a clear message
4. Preserves all completed work

### To resume after the quota resets

```bash
# 1. Remove the quota marker
rm serge/eval/results/_quota_hit

# 2. Re-run — picks up exactly where it left off
./eval/run.sh
```

The pipeline will skip everything already completed and continue with pending items.

Alternatively, bypass the marker check:
```bash
./eval/run.sh --force-retry  # overrides marker too
```

## Debugging failures

### See which runs failed
```bash
cat serge/eval/results/_failures.log
```
Each failure has a timestamp, scenario ID, attempt number, and classified error (`quota`, `network`, `unknown`).

### Re-run only failed items
Failed items have invalid or missing output files, so they are **pending** by default.
A plain `./eval/run.sh` will retry them automatically.

### Inspect a specific run
```bash
# Agent answer
cat serge/eval/results/L4-01/run_1.json | jq .

# Judge score
cat serge/eval/results/L4-01/run_1_judge.json | jq .
```

## Configuration

Edit `eval/config.json` to change:
- `runs_per_scenario` (default: 3)
- `model` (default: `sonnet`) — the agent model
- `judge_model` (default: `sonnet`) — the judge model
- `max_budget_per_run_usd` (default: 0.50) — hard limit per agent call

Retry and backoff settings live at the top of `eval/run.sh`:
- `MAX_RETRIES=3`
- `BACKOFF_BASE=15` seconds (exponential: 15, 45, 135)

## Safe interrupt

Press `Ctrl+C` any time. The script traps the signal, writes the message `Interrupted. Progress is preserved — re-run the script to resume.`, and exits with code 130.

All completed runs are saved on disk. Re-run `./eval/run.sh` to pick up where you stopped.

## File layout

```
eval/
├── run.sh                        # the pipeline
├── config.json                   # models, runs per scenario, budget
├── judge_prompt.md               # instructions for the judge LLM
├── judge_schema.json             # expected judge output shape
├── scenarios/                    # 27 scenario definitions (input)
│   ├── L1-01-product-lookup.json
│   └── ...
├── results/                      # per-run output (generated)
│   ├── L1-01/
│   │   ├── run_1.json            # agent response
│   │   └── run_1_judge.json      # judge score
│   ├── _failures.log             # log of failed attempts
│   └── _quota_hit                # marker (exists only after a quota hit)
└── report/
    └── evaluation_report.json    # aggregated report (generated)
```
