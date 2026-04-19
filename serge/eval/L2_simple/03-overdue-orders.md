# Overdue Orders

## Metadata
- **Level:** L2
- **Use Case:** UC-4
- **Systems:** OMS
- **Expected tools:** 1
- **Realistic user:** Ops manager doing a daily check

## Question
"Which customer orders are currently overdue?"

## Required Facts

**Critical:**
- ORD-007 must be listed (processing, required Feb 12, well overdue)
- ORD-011 must be listed (shipped but stuck in exception)

**Important:**
- Each overdue order includes: ID, customer, required date, current status

**Nice to have:**
- Sorted by severity (most overdue first)

## Forbidden Content
- Orders that are delivered or cancelled cannot appear
- Future-required orders (not yet overdue) must not be listed as overdue

## Flexibility Notes
- If the agent decides to do follow-up investigation (e.g., why ORD-007 is overdue), that's a bonus
- The evaluator should NOT penalize brevity — a simple list is a valid answer at this level

## Ground Truth

Key scenario orders that must appear:
| Order | Status | Required Date |
|-------|--------|---------------|
| ORD-007 | processing | 2026-02-12 |
| ORD-011 | shipped (exception) | 2026-02-10 |
