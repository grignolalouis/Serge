# Operational Overview — All Current Problems

## Use Case
UC-2: Root Cause Analysis | Difficulty: hard

## Question
"Give me a summary of all current operational problems in our supply chain."

## Expected Tool Calls
1. `oms_list_overdue_orders()` — find orders past due
2. `tms_list_exception_shipments()` — find shipment problems
3. `wms_list_low_stock()` — find inventory alerts
4. `srm_list_overdue_purchase_orders()` — find late supplier deliveries

## Ground Truth
**Overdue orders** (past required date, not delivered/cancelled):
- ORD-007: processing, urgent, required Feb 12, CUST-001 (Siam Paragon) — 50x PROD-001
- ORD-011: shipped, required Feb 10, CUST-003 (Tops) — shipment has exception
- ORD-009: pending, required Feb 18, CUST-008 (Gaggan) — may be overdue depending on current date

**Shipment exceptions:**
- SHP-009: ORD-011, refrigeration malfunction, Kerry Express

**Low stock alerts:**
- PROD-001 (Nam Doc Mai): 8 available, reorder point 20 — critically low
- PROD-005 (Mixed Assortment): 12 available, reorder point 15 — low

**Overdue POs:**
- Depends on current date vs expected delivery dates. PO-008 (expected Feb 18), PO-009 (expected Feb 15) may be overdue.

## Must NOT Contain
- Fabricated problems not in the data
- Missed known issues

## Evaluation Checklist
- [ ] All 4 systems queried
- [ ] Overdue orders identified
- [ ] Exception shipment identified (SHP-009)
- [ ] Low stock products identified (PROD-001, PROD-005)
- [ ] Connections made (e.g., ORD-007 stuck because PROD-001 is low)
- [ ] Structured, actionable summary
