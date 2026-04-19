# Customer Impact Assessment

## Metadata
- **Level:** L4
- **Use Case:** UC-2
- **Systems:** WMS, OMS
- **Expected tools:** 3–5
- **Realistic user:** Ops manager preparing customer communications

## Question
"Given our PROD-001 shortage, which customers are at risk of being impacted?"

## Required Facts

**Critical:**
- Current PROD-001 stock: 8 units
- ORD-007 (CUST-001 Siam Paragon) needs 50 units — **blocked** (urgent, already overdue)
- ORD-009 (CUST-008 Gaggan Anand) needs 5 units — at risk (high priority, required Feb 18)
- ORD-023 may also be affected (CUST-004 Mandarin Oriental, 10 units, shipped — already accounted for)

**Important:**
- Identify the customer names (not just IDs) from OMS
- Quantify the gap: 50 + 5 = 55 units needed vs 8 on hand

**Nice to have:**
- Mention PO-008 timing (arrives Feb 18) and which orders it rescues
- Prioritization suggestion

## Forbidden Content
- Customers whose orders don't include PROD-001
- Already-delivered orders listed as "at risk"

## Flexibility Notes
- Agent may use `oms_search_customer_orders_by_product` to get the list then filter to active ones
- Distinguishing active demand (pending/processing) from historical is critical

## Ground Truth

Active (non-delivered) PROD-001 demand from scenario data:
| Order | Customer | Qty | Status | Priority |
|-------|----------|-----|--------|----------|
| ORD-007 | CUST-001 Siam Paragon | 50 | processing | urgent |
| ORD-009 | CUST-008 Gaggan Anand | 5 | pending | high |
| ORD-023 | CUST-004 Mandarin Oriental | 10 | shipped | urgent |

Total active demand: **65 units**; available: **8**; shortfall: **~57**.
