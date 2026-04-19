# Carrier Mismatch Diagnosis

## Metadata
- **Level:** L4
- **Use Case:** UC-2
- **Systems:** TMS, OMS, WMS
- **Expected tools:** 4–5
- **Realistic user:** Ops manager investigating the SHP-009 incident

## Question
"Shipment SHP-009 had a refrigeration failure. Was choosing Kerry Express a mistake given the product?"

## Required Facts

**Critical:**
- SHP-009: Kerry Express (CAR-001, type: ground)
- Associated order: ORD-011, product PROD-003 (Keaw Mango, storage_type: refrigerated)
- **Yes, mismatch**: ground carrier shouldn't carry refrigerated product
- At least one refrigerated carrier exists (SCG Cold Chain or Flash Express)

**Important:**
- Reason recorded: "Refrigeration unit malfunction on delivery vehicle"
- Affected customer: Tops Supermarket (CUST-003)
- Order is now overdue

**Nice to have:**
- Recommendation of appropriate carrier (SCG Cold Chain reasonable at 25 THB/kg vs 15 THB/kg)
- Cost comparison

## Forbidden Content
- Claim Kerry Express was the correct choice
- Missing the storage_type mismatch

## Flexibility Notes
- Multiple valid tool paths: shipment → order → product, or start from the product
- The judge checks that the agent makes the connection (refrigerated product + ground carrier = problem)

## Ground Truth

| Key | Value |
|-----|-------|
| shipment_id | SHP-009 |
| carrier | Kerry Express (CAR-001, type: ground) |
| order_id | ORD-011 |
| product | PROD-003 (Keaw Mango Box 5kg) |
| storage_type | refrigerated |
| verdict | Mismatch — refrigerated product should have used CAR-002 or CAR-003 |
| exception_reason | Refrigeration unit malfunction on delivery vehicle |
