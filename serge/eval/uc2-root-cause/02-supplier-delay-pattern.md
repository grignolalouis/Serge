# Supplier Delay Pattern — Repeated Late Deliveries

## Use Case
UC-2: Root Cause Analysis | Difficulty: medium

## Question
"Why was purchase order PO-002 late, and is this a recurring problem with that supplier?"

## Expected Tool Calls
1. `srm_get_purchase_order(purchase_order_id="PO-002")` — see it was late
2. `srm_get_supplier(supplier_id="SUP-002")` — get supplier profile
3. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-002")` — check all POs from this supplier

## Ground Truth
- PO-002: SUP-002 (Northern Fruits Co.), PROD-002 (Ok Rong 10kg), 80 units @ 400 THB
- Ordered 2026-01-05, expected 2026-01-10, received **2026-01-13** — **3 days late**
- SUP-002 profile: Chiang Mai, rating 3.2, declared lead time 5 days
- **Pattern**: All received POs from SUP-002 are late:
  - PO-002: expected Jan 10, received Jan 13 (3 days late)
  - PO-005: expected Jan 20, received Jan 24 (4 days late)
  - PO-009: pending (ordered Feb 10, expected Feb 15)
- SUP-002 on-time rate: **0%** (0 out of 2 received POs on time)
- Average actual lead time: ~8 days (vs declared 5)

## Must NOT Contain
- Claim that SUP-002 has good delivery performance
- Incorrect PO dates

## Evaluation Checklist
- [ ] PO-002 delay correctly identified (3 days)
- [ ] Pattern recognized across multiple POs
- [ ] Both PO-002 and PO-005 cited as late
- [ ] 0% on-time rate mentioned or calculated
- [ ] Supplier rating (3.2) and declared lead time (5 days) noted
- [ ] No hallucinated POs
