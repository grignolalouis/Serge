# SERGE Evaluation — Scenarios by Complexity

27 scenarios organized by **resolution complexity** (L1 trivial → L5 expert), covering the full capability spectrum of the agent.

## Complexity Levels

| Level | Name | Systems | Tools | What it tests |
|-------|------|---------|-------|---------------|
| **L1** | Trivial | 1 | 1 | Single lookup — can the agent retrieve one entity? |
| **L2** | Simple | 1 | 2–3 | Single-system query — filtering, listing |
| **L3** | Moderate | 2 | 3–5 | Cross-reference between two systems |
| **L4** | Complex | 3–4 | 5–8 | Causal reasoning across multiple systems |
| **L5** | Expert | 3–4 | 8+ | Full diagnostic + actionable recommendations |

## Use Cases (orthogonal to complexity)

Each scenario is tagged with its business use case:

- **UC-1 Order Status** — What's happening with a specific order?
- **UC-2 Root Cause** — Why is something not working as expected?
- **UC-3 Supplier/Sourcing** — Compare and evaluate supply options
- **UC-4 Operational** — Discover and monitor overall state
- **UC-5 Planning** — Can we do X, and how?

## Evaluation Philosophy

The judge must be **flexible on form, strict on facts**.

### What's flexible
- Tool call order and selection (multiple valid paths lead to the same answer)
- Wording ("450 THB" = "฿450" = "450 baht")
- Level of detail (brief summary vs exhaustive report are both valid)
- Extra analysis beyond the minimum (bonus, not required)
- Choice to batch vs sequence queries

### What's strict
- **Factual accuracy** — numbers, IDs, dates, statuses must match ground truth
- **No hallucinations** — nothing that isn't in the seed data
- **Coherent reasoning** — the answer must be internally consistent

### Scoring (5-point scale)

| Dimension | Weight | Tolerance |
|-----------|--------|-----------|
| Factual correctness | 40% | Strict — wrong values fail |
| Completeness (critical facts) | 30% | Strict — missing core facts fail |
| No hallucination | 20% | Strict — fabricated data fails |
| Reasoning clarity | 10% | Lenient — any coherent explanation passes |

Tool count and response time are measured but do NOT affect the score.

## Scenario Format

See `scenario-template.md`. Every scenario defines:

1. **Metadata** — level, use case, systems, expected tool count range
2. **Question** — realistic ops manager phrasing
3. **Required Facts** — critical / important / nice-to-have tiers
4. **Forbidden Content** — hallucinations or errors that must NOT appear
5. **Flexibility Notes** — what the evaluator should and should not penalize
6. **Ground Truth** — reference values the judge can verify against

## Pass Threshold

- **Pass:** ≥ 3.5/5.0 weighted score, all critical facts present, no forbidden content
- **Strong pass:** ≥ 4.0/5.0
- **Full score:** 5.0 (rare — reserved for exceptional answers)

## Running the Evaluation

Each scenario is run independently. The judge compares the agent's full response (not just tool calls) against the ground truth. The agent's internal tool trace is used as auxiliary evidence but not directly scored.
