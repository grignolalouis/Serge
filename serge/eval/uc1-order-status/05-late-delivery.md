# Late Delivery

## Use Case
UC-1: Order Status Inquiry | Difficulty: medium

## Question
"Was order ORD-003 delivered on time?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-003")` — get order details with required date
2. `tms_track_order(order_id="ORD-003")` — get shipment with delivery date

## Ground Truth
- Order ORD-003: customer CUST-003 (Tops Supermarket), status **delivered**
- Ordered 2026-01-15, required by **2026-01-18**, delivered **2026-01-20**
- **2 days late** (delivered Jan 20, required Jan 18)
- Lines: 20x PROD-002 @ 650 THB, 15x PROD-003 @ 350 THB = 18,250 THB
- Shipment SHP-003: carrier Kerry Express (CAR-001, ground), 275kg
- Shipped 2026-01-17, ETA 2026-01-19, delivered 2026-01-20
- Shipment itself was also late vs ETA (ETA Jan 19, delivered Jan 20)

## Must NOT Contain
- Claim that delivery was on time
- Wrong delivery date

## Evaluation Checklist
- [ ] Correct tools called
- [ ] Late delivery clearly identified
- [ ] Specific dates compared (required vs delivered)
- [ ] Days late calculated correctly (2 days)
- [ ] Carrier identified (Kerry Express ground)
- [ ] No hallucinated data
