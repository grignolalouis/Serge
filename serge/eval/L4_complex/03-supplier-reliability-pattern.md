# Supplier Reliability Pattern

## Metadata
- **Level:** L4
- **Use Case:** UC-3
- **Systems:** SRM
- **Expected tools:** 2–4
- **Realistic user:** Procurement deciding whether to continue with a supplier

## Question
"Does Northern Fruits Co. (SUP-002) have a chronic delivery problem, or is it an isolated issue?"

## Required Facts

**Critical:**
- SUP-002 has **repeated** late deliveries in scenario data
- PO-002: expected Jan 10, received Jan 13 (3 days late)
- PO-005: expected Jan 20, received Jan 24 (4 days late)
- Pattern is **chronic** — 0% on-time rate for received POs
- PO-009 (pending) and PO-018 (cancelled) are also with this supplier

**Important:**
- Average actual lead time (~8 days) exceeds declared (5 days) by ~60%
- Rating 3.2 is consistent with these metrics

**Nice to have:**
- Recommendation (diversify sourcing, renegotiate, or replace)
- Comparison to better-performing suppliers

## Forbidden Content
- Claim delivery is on time
- Invented POs

## Flexibility Notes
- Agent may use `srm_compare_suppliers` as a shortcut or drill into the specific supplier
- The conclusion "chronic problem" is the core ask; how the agent arrives at it is flexible

## Ground Truth

Scenario POs for SUP-002:
| PO | Expected | Actual | Days Late |
|----|----------|--------|-----------|
| PO-002 | 2026-01-10 | 2026-01-13 | 3 |
| PO-005 | 2026-01-20 | 2026-01-24 | 4 |
| PO-009 | 2026-02-15 | pending | n/a |
| PO-018 | 2026-02-02 | cancelled | n/a |

On-time rate (scenario received POs): **0/2 = 0%**
