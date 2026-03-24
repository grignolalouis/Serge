# Carrier Exception — Refrigeration Failure Investigation

## Use Case
UC-2: Root Cause Analysis | Difficulty: medium

## Question
"What caused the delivery failure for shipment SHP-009 and what's the impact?"

## Expected Tool Calls
1. `tms_get_shipment(shipment_id="SHP-009")` — get shipment details + tracking
2. `oms_get_order(order_id="ORD-011")` — get the affected order
3. `oms_get_customer(customer_id="CUST-003")` — identify affected customer
4. `tms_get_carrier(carrier_id="CAR-001")` — check carrier type

## Ground Truth
- SHP-009: order ORD-011, carrier Kerry Express (CAR-001), status **exception**
- Exception: "Refrigeration unit malfunction on delivery vehicle"
- Carrier CAR-001 is type **ground** (not refrigerated) — potential carrier selection issue
- Order ORD-011: CUST-003 (Tops Supermarket), 25x PROD-003 (Keaw Mango 5kg) @ 350 THB = 8,750 THB
- Required by 2026-02-10 — **overdue**
- Tracking: picked up Feb 6 → in transit Feb 6-7 → exception Feb 8 at Pathum Thani
- Note: Kerry Express is a "ground" carrier but shipment had refrigeration — possible mismatch in data, or Kerry subcontracts refrigerated vehicles

## Must NOT Contain
- Fabricated resolution date
- Wrong exception reason

## Evaluation Checklist
- [ ] Exception reason correctly identified (refrigeration malfunction)
- [ ] Affected order and customer identified
- [ ] Overdue status noted
- [ ] Carrier type noted (ground)
- [ ] Tracking timeline referenced
- [ ] Impact assessed (customer waiting, perishable goods at risk)
