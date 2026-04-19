# Supplier Order History

## Metadata
- **Level:** L2
- **Use Case:** UC-3
- **Systems:** SRM
- **Expected tools:** 1
- **Realistic user:** Procurement reviewing a supplier relationship

## Question
"Show me all our purchase orders from Somchai Mango Farm (SUP-001)."

## Required Facts

**Critical:**
- 4 scenario POs: PO-001, PO-004, PO-006, PO-008
- Each with status, expected delivery, and actual delivery (if received)

**Important:**
- PO-001 received on time (Jan 5)
- PO-004 received 1 day early (Jan 14 vs expected Jan 15)
- PO-006 received on time (Jan 25)
- PO-008 status: shipped, expected Feb 18

## Forbidden Content
- POs from other suppliers
- Fabricated PO IDs

## Flexibility Notes
- Agent may also pull additional generated POs (ID >= 100) — include is OK, omit is OK
- Chronological ordering is natural but not required

## Ground Truth

Scenario POs from SUP-001:
| PO | Status | Ordered | Expected | Actual |
|----|--------|---------|----------|--------|
| PO-001 | received | 2026-01-02 | 2026-01-05 | 2026-01-05 |
| PO-004 | received | 2026-01-12 | 2026-01-15 | 2026-01-14 |
| PO-006 | received | 2026-01-22 | 2026-01-25 | 2026-01-25 |
| PO-008 | shipped | 2026-02-10 | 2026-02-18 | — |
