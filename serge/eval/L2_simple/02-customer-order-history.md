# Customer Order History

## Metadata
- **Level:** L2
- **Use Case:** UC-4
- **Systems:** OMS
- **Expected tools:** 1–2
- **Realistic user:** Account manager reviewing a client's activity

## Question
"Show me all the orders from Mandarin Oriental (CUST-004)."

## Required Facts

**Critical:**
- Multiple orders are listed (at least the scenario ones: ORD-004, ORD-012, ORD-023)
- Each order has a status

**Important:**
- ORD-004 is delivered
- ORD-012 is processing
- ORD-023 is shipped

**Nice to have:**
- Dates for each order
- Comment on the pattern (regular customer)

## Forbidden Content
- Orders for other customers
- Invented order IDs

## Flexibility Notes
- The dataset has many generated orders (IDs ≥ 100) for CUST-004 — the agent should include them all, but the judge should focus on the scenario orders being correctly identified
- Any order format (table, bullets) is fine

## Ground Truth

Scenario orders for CUST-004:
| Order | Status | Ordered | Required |
|-------|--------|---------|----------|
| ORD-004 | delivered | 2026-01-19 | 2026-01-24 |
| ORD-012 | processing | 2026-02-11 | 2026-02-17 |
| ORD-023 | shipped | 2026-02-10 | 2026-02-15 |
