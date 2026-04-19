# Fulfillment Feasibility Plan

## Metadata
- **Level:** L5
- **Use Case:** UC-5 (planning)
- **Systems:** OMS, WMS, SRM, TMS
- **Expected tools:** 6–9
- **Realistic user:** Ops manager under customer pressure

## Question
"Order ORD-028 needs 12 units of PROD-001 and 6 units of PROD-004 for The Sukhothai Hotel by Feb 18. Can we fulfill it? Check stock, identify shortages, find incoming POs, and recommend a carrier — both products need refrigeration."

## Required Facts

**Critical:**
- PROD-004 (19 available): **covered** for 6 units
- PROD-001 (8 available): **short by 4 units** for the 12 needed
- PO-008 brings 150× PROD-001 on Feb 18 — same day as deadline, tight
- Both products are refrigerated → use CAR-002 (SCG Cold Chain) or CAR-003 (Flash Express), NOT Kerry Express

**Important:**
- Total weight: 12×5 + 6×10 = 120 kg
- Carrier cost estimates (SCG: 120×25 = 3,000 THB; Flash: 120×35 = 4,200 THB)
- Recommendation: SCG Cold Chain (same transit time, cheaper)

**Nice to have:**
- Fulfillment options: wait for PO-008 (risky), partial ship now + complete later, reallocate from PO-008 when it arrives
- Note that ORD-007 also competes for PROD-001 stock
- Mention of shelf-life risk for PROD-004 (5 days)

## Forbidden Content
- Claim that stock fully covers the order (it doesn't — 4-unit gap on PROD-001)
- Recommend a ground carrier (type mismatch)
- Fabricated PO that resolves PROD-001 shortage earlier than PO-008

## Flexibility Notes
- This scenario tests planning + synthesis, not a narrow factual answer
- Multiple valid conclusions: partial shipment, rely on PO-008, escalate — all defensible
- The judge should look for: factually correct diagnosis + coherent plan, not a specific plan

## Ground Truth

| Key | Value |
|-----|-------|
| ORD-028 customer | CUST-006 (The Sukhothai Hotel, wholesale) |
| ORD-028 required | 2026-02-18 |
| PROD-001 available | 8 (need 12, short 4) |
| PROD-004 available | 19 (need 6, covered) |
| Relevant incoming PO | PO-008 (150× PROD-001, ETA 2026-02-18) |
| Carrier options | SCG Cold Chain (25 THB/kg), Flash Express (35 THB/kg) |
| Carrier NOT to use | Kerry Express (ground — mismatch) |
| Total weight | 120 kg |
