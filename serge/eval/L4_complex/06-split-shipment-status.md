# Split Shipment Status

## Metadata
- **Level:** L4
- **Use Case:** UC-1
- **Systems:** OMS, TMS, WMS
- **Expected tools:** 3–5
- **Realistic user:** CS or ops tracking a complex order

## Question
"Order ORD-023 was split into two shipments. Give me the status of each and explain why it was split."

## Required Facts

**Critical:**
- ORD-023 has **two shipments**: SHP-017 and SHP-018
- SHP-017 via refrigerated carrier (SCG Cold Chain, CAR-002) — for PROD-001 (refrigerated)
- SHP-018 via ground carrier (Kerry Express, CAR-001) — for PROD-006 (ambient)
- Both currently in transit

**Important:**
- The split is justified: PROD-001 needs cold chain, PROD-006 (dried mango) doesn't
- The order note explicitly mentions "Split into 2 shipments — fresh and ambient"
- Customer: Mandarin Oriental (CUST-004)

**Nice to have:**
- Comment that this is the correct carrier strategy per product storage_type
- ETAs for each shipment

## Forbidden Content
- Missing either shipment
- Wrong carrier assignment
- Claim that the split was an error

## Flexibility Notes
- `tms_track_customer_order` returns both shipments as a list (new split-shipment support)
- Agent may also fetch product details to explain the reason; bonus

## Ground Truth

| Shipment | Carrier | Type | Weight | Content rationale |
|----------|---------|------|--------|-------------------|
| SHP-017 | SCG Cold Chain (CAR-002) | refrigerated | 50 kg | PROD-001 (fresh, needs cold chain) |
| SHP-018 | Kerry Express (CAR-001) | ground | 30 kg | PROD-006 (dried, ambient-safe) |

Customer: **CUST-004 Mandarin Oriental Bangkok**.
Split reason: storage type difference between products.
