# Order Tracking with Current Location

## Metadata
- **Level:** L3
- **Use Case:** UC-1
- **Systems:** OMS, TMS
- **Expected tools:** 2–3
- **Realistic user:** CS answering a "where is my package" question

## Question
"Where is order ORD-005 right now?"

## Required Facts

**Critical:**
- Status: shipped / in transit
- Carrier: SCG Cold Chain (CAR-002, refrigerated)
- Tracking number: SCG-20260210-005
- ETA: Feb 12, 2026

**Important:**
- Shipped date: Feb 10, 2026
- Destination: Villa Market, Bangkok
- Most recent tracking event: at local distribution center, Feb 11

**Nice to have:**
- Full tracking event timeline
- Confidence about on-time delivery

## Forbidden Content
- Claim that the order is already delivered
- Fabricated tracking events

## Flexibility Notes
- Agent may use `tms_track_customer_order` (which wraps shipment + events) or combine `oms_get_customer_order` + `tms_get_shipment`
- Both approaches are equally valid

## Ground Truth

| Key | Value |
|-----|-------|
| order_id | ORD-005 |
| shipment_id | SHP-005 |
| carrier | SCG Cold Chain (CAR-002) |
| status | in_transit |
| tracking_number | SCG-20260210-005 |
| shipped_at | 2026-02-10 08:00 |
| eta | 2026-02-12 17:00 |
| last_event | "At local distribution center" (Feb 11, 10:00) |
