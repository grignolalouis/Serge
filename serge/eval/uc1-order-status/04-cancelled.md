# Cancelled Order

## Use Case
UC-1: Order Status Inquiry | Difficulty: easy

## Question
"Can you check on order ORD-014?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-014")` — get order details
2. `tms_track_order(order_id="ORD-014")` — check for shipment (should find none)

## Ground Truth
- Order ORD-014: customer CUST-005 (Villa Market), status **cancelled**
- Ordered 2026-01-22, required by 2026-01-28
- Priority: normal
- Lines: 20x PROD-002 (Ok Rong 10kg) @ 650 THB = 13,000 THB
- No shipment exists (tms_track_order returns error)
- No shipped_date, no delivered_date

## Must NOT Contain
- Shipment details (none exist)
- Status reported as anything other than cancelled
- Fabricated cancellation reason (not in seed data)

## Evaluation Checklist
- [ ] Correct tools called
- [ ] Status correctly identified as cancelled
- [ ] No shipment correctly noted
- [ ] Order details accurate
- [ ] No fabricated cancellation reason
