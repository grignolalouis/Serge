# Happy Path — Delivered Order

## Use Case
UC-1: Order Status Inquiry | Difficulty: easy

## Question
"What is the status of order ORD-001?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-001")` — get order details
2. `tms_track_order(order_id="ORD-001")` — get shipment and tracking

## Ground Truth
- Order ORD-001: customer CUST-001 (Siam Paragon Food Hall), status **delivered**
- Ordered 2026-01-08, required by 2026-01-14, shipped 2026-01-10, delivered 2026-01-11
- Priority: normal
- Lines: 10x PROD-001 (Nam Doc Mai 5kg) @ 450 THB, 5x PROD-003 (Keaw 5kg) @ 350 THB
- Total: 6,250 THB
- Shipment SHP-001: carrier SCG Cold Chain (CAR-002), tracking SCG-20260110-001
- Shipped 2026-01-10 08:00, ETA 2026-01-11 17:00, delivered 2026-01-11 14:30
- Delivered **on time** (3 days before required date)

## Must NOT Contain
- Wrong carrier name
- Incorrect quantities or prices
- Fabricated tracking events (SHP-001 has no tracking events in seed data)

## Evaluation Checklist
- [ ] Correct tools called (oms_get_order + tms_track_order)
- [ ] Order status correctly identified as delivered
- [ ] Line items and quantities accurate
- [ ] Shipment details accurate (carrier, dates)
- [ ] On-time delivery noted
- [ ] No hallucinated data
