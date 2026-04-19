# Late Delivery Check

## Metadata
- **Level:** L3
- **Use Case:** UC-1
- **Systems:** OMS, TMS
- **Expected tools:** 2
- **Realistic user:** CS investigating a customer complaint

## Question
"Was order ORD-003 delivered on time?"

## Required Facts

**Critical:**
- Required date: Jan 18, 2026
- Delivered date: Jan 20, 2026
- **2 days late**

**Important:**
- Customer: Tops Supermarket (CUST-003)
- Carrier: Kerry Express (ground)
- Shipped Jan 17, ETA Jan 19, delivered Jan 20

**Nice to have:**
- Observation that shipment also missed its ETA by 1 day

## Forbidden Content
- Claim that delivery was on time
- Wrong dates

## Flexibility Notes
- Two-step retrieval: order first, then shipment (or track_customer_order)
- The arithmetic (2 days late) should be explicit or clearly implied

## Ground Truth

| Key | Value |
|-----|-------|
| order_id | ORD-003 |
| required_date | 2026-01-18 |
| delivered_date | 2026-01-20 |
| days_late | 2 |
| shipment_id | SHP-003 |
| carrier | Kerry Express (CAR-001) |
