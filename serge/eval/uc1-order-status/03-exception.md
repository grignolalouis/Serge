# Exception — Carrier Delivery Failure

## Use Case
UC-1: Order Status Inquiry | Difficulty: medium

## Question
"What's happening with order ORD-011? The customer is asking."

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-011")` — get order details
2. `tms_track_order(order_id="ORD-011")` — get shipment and tracking events

## Ground Truth
- Order ORD-011: customer CUST-003 (Tops Supermarket), status **shipped**
- Ordered 2026-02-04, required by 2026-02-10, shipped 2026-02-06
- Lines: 25x PROD-003 (Keaw 5kg) @ 350 THB = 8,750 THB
- Shipment SHP-009: carrier Kerry Express (CAR-001), status **exception**
- Exception reason: "Refrigeration unit malfunction on delivery vehicle"
- Tracking events:
  1. Feb 6 09:00 — Bang Na — picked up
  2. Feb 6 15:00 — Bangkok — departed sorting facility
  3. Feb 7 11:00 — Nonthaburi — in transit
  4. Feb 8 08:00 — Pathum Thani — **exception**: refrigeration unit malfunction
- Order is **overdue** (required Feb 10, not delivered)

## Must NOT Contain
- Status reported as delivered
- Wrong exception reason
- Fabricated resolution or new delivery date

## Evaluation Checklist
- [ ] Correct tools called
- [ ] Exception status clearly communicated
- [ ] Exception reason mentioned (refrigeration malfunction)
- [ ] Overdue status noted (past required date)
- [ ] Tracking timeline presented
- [ ] No hallucinated resolution
