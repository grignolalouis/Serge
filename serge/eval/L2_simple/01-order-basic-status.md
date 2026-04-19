# Basic Order Status

## Metadata
- **Level:** L2
- **Use Case:** UC-1
- **Systems:** OMS
- **Expected tools:** 1–2
- **Realistic user:** CS rep checking on an order

## Question
"What's the current status of order ORD-005?"

## Required Facts

**Critical:**
- Status: shipped
- Customer: CUST-005 (Villa Market)
- Ordered: Feb 8, 2026
- Required by: Feb 14, 2026
- Shipped: Feb 10, 2026

**Important:**
- Line items: 15× PROD-002 and 10× PROD-005
- Priority: normal

**Nice to have:**
- Comment that the order is ahead of required date

## Forbidden Content
- Claim that the order is delivered or cancelled
- Wrong customer
- Fabricated tracking information (unless TMS is queried too)

## Flexibility Notes
- Agent may also call `tms_track_customer_order` to enrich the answer (bonus)
- Agent may or may not compute the total amount

## Ground Truth

| Key | Value |
|-----|-------|
| id | ORD-005 |
| status | shipped |
| customer_id | CUST-005 |
| order_date | 2026-02-08 |
| required_date | 2026-02-14 |
| shipped_date | 2026-02-10 |
| priority | normal |
