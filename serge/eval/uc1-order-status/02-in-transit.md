# In Transit — Active Shipment with Tracking

## Use Case
UC-1: Order Status Inquiry | Difficulty: easy

## Question
"Where is order ORD-005 right now?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-005")` — get order details
2. `tms_track_order(order_id="ORD-005")` — get shipment and tracking events

## Ground Truth
- Order ORD-005: customer CUST-005 (Villa Market), status **shipped**
- Ordered 2026-02-08, required by 2026-02-14, shipped 2026-02-10
- Priority: normal
- Lines: 15x PROD-002 (Ok Rong 10kg) @ 650 THB, 10x PROD-005 (Mixed 5kg) @ 500 THB
- Total: 14,750 THB
- Shipment SHP-005: carrier SCG Cold Chain (CAR-002), status **in_transit**
- Tracking SCG-20260210-005, weight 200kg
- Shipped 2026-02-10 08:00, ETA 2026-02-12 17:00
- Tracking events:
  1. Feb 10 08:30 — Bang Na, Bangkok — picked up
  2. Feb 10 14:00 — Bangkok — departed sorting facility
  3. Feb 11 10:00 — Bangkok — at local distribution center, scheduled for delivery

## Must NOT Contain
- Status reported as delivered
- Fabricated delivery date
- Wrong carrier (must be SCG Cold Chain, not Kerry Express)

## Evaluation Checklist
- [ ] Correct tools called
- [ ] Status correctly identified as in transit
- [ ] Tracking events listed in order
- [ ] ETA mentioned (Feb 12)
- [ ] Carrier correctly identified as SCG Cold Chain
- [ ] No hallucinated data
