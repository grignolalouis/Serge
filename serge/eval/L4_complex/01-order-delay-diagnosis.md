# Order Delay Diagnosis

## Metadata
- **Level:** L4
- **Use Case:** UC-2
- **Systems:** OMS, TMS, WMS, SRM
- **Expected tools:** 4–6
- **Realistic user:** Ops manager explaining a delay to a customer

## Question
"Why hasn't order ORD-007 shipped yet? The customer is chasing us."

## Required Facts

**Critical:**
- ORD-007 is in "processing" status, past its required date (Feb 12)
- It needs 50 units of PROD-001 (Nam Doc Mai)
- Current inventory: only 8 units available
- Root cause: **stock shortage**
- A supplier order (PO-008) is in transit, expected Feb 18 — will arrive AFTER required date

**Important:**
- Customer: Siam Paragon (CUST-001), priority urgent
- PO-008 is from SUP-001 (Somchai Mango Farm), 150 units
- No shipment exists yet for ORD-007

**Nice to have:**
- Recommendation (partial fulfillment? communicate with customer?)
- Mention of competing demand (ORD-009 also needs PROD-001)

## Forbidden Content
- Wrong root cause (e.g., claiming carrier issue or customer cancellation)
- Claim that PO-008 will arrive before Feb 12
- Fabricated IDs

## Flexibility Notes
- The canonical path is OMS → TMS → WMS → SRM, but any path that surfaces the 4 critical facts is acceptable
- The agent may use `oms_search_customer_orders_by_product` as an alternative entry point — also valid
- Adding proposed actions is bonus, not required

## Ground Truth

| Key | Value |
|-----|-------|
| order_id | ORD-007 |
| status | processing |
| required_date | 2026-02-12 |
| customer | CUST-001 (Siam Paragon) |
| needs | 50× PROD-001 |
| available | 8 |
| incoming_po | PO-008 (150 units, shipped, ETA 2026-02-18) |
| root_cause | stock shortage |
