# Full Operational Health Check

## Metadata
- **Level:** L5
- **Use Case:** UC-4
- **Systems:** OMS, TMS, WMS, SRM
- **Expected tools:** 6–10
- **Realistic user:** COO wanting a Monday-morning state-of-the-business view

## Question
"Give me a full health check of our supply chain: orders, shipments, inventory, supplier performance. Where are the risks and what should I act on this week?"

## Required Facts

**Critical:**
The answer must surface issues across all four systems:
- **OMS:** Overdue orders (ORD-007, ORD-011)
- **TMS:** Shipment exceptions (SHP-009)
- **WMS:** Low stock (PROD-001, PROD-005, PROD-008)
- **SRM:** Supplier reliability concern (SUP-002 chronic delays)

**Important:**
- At least one **cross-system insight** (e.g., ORD-007 stuck because PROD-001 is low; SHP-009 failed because of carrier mismatch)
- Concrete prioritized actions, not just a list

**Nice to have:**
- Coverage of positives (which suppliers are reliable, which orders are on track)
- Metrics or counts

## Forbidden Content
- Fabricated issues
- Missing major known problems

## Flexibility Notes
- The agent will use many tools (probably 6–10). The judge should NOT count tools — the quality of synthesis is what matters.
- Different agents may prioritise items differently. Any reasonable prioritization passes.
- Length can vary from a concise executive summary to a detailed report; both are acceptable if they cover the critical facts.

## Ground Truth (scenario snapshot)

| System | Issue | Details |
|--------|-------|---------|
| OMS | ORD-007 overdue | processing, urgent, required Feb 12 |
| OMS | ORD-011 overdue | shipped but stuck in exception |
| TMS | SHP-009 exception | refrigeration malfunction, Kerry Express |
| WMS | PROD-001 low | 8 / 20 reorder point |
| WMS | PROD-005 low | 12 / 15 |
| WMS | PROD-008 low | 6 / 15 |
| SRM | SUP-002 chronic | 0% on-time in scenario data |

Key cross-links:
- ORD-007 ↔ PROD-001 low stock
- SHP-009 ↔ carrier type mismatch with PROD-003 storage
