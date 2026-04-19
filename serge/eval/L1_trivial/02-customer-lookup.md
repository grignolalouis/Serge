# Customer Lookup

## Metadata
- **Level:** L1
- **Use Case:** UC-4
- **Systems:** OMS
- **Expected tools:** 1
- **Realistic user:** Sales or ops looking up a client

## Question
"Who is customer CUST-004?"

## Required Facts

**Critical:**
- Company: Mandarin Oriental Bangkok
- Segment: wholesale

**Important:**
- Contact: Arunee Thongchai
- Region: Bangkok
- Delivery zone: central_bkk

**Nice to have:**
- Address

## Forbidden Content
- Wrong company name or segment
- Fabricated order history (unless another tool is called)

## Flexibility Notes
- Agent may stop at the customer profile; it does not need to list orders

## Ground Truth

| Key | Value |
|-----|-------|
| id | CUST-004 |
| name | Arunee Thongchai |
| company | Mandarin Oriental Bangkok |
| segment | wholesale |
| region | Bangkok |
| delivery_zone | central_bkk |
