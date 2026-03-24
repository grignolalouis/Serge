# Order Stuck in Processing — Is Stock the Issue?

## Use Case
UC-2: Root Cause Analysis | Difficulty: medium

## Question
"Order ORD-012 has been in processing for a while. What's holding it up?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-012")` — see order details
2. `wms_check_inventory(product_id="PROD-002")` — check stock for the ordered product
3. `tms_track_order(order_id="ORD-012")` — confirm no shipment yet

## Ground Truth
- ORD-012: CUST-004 (Mandarin Oriental), status **processing**, priority **high**
- Ordered 2026-02-11, required by 2026-02-17
- Lines: 8x PROD-002 (Ok Rong 10kg) @ 650 THB = 5,200 THB
- WMS: PROD-002 has 45 qty, 10 reserved = **35 available** — stock is **sufficient** (needs only 8)
- No shipment created yet
- **Analysis**: Stock is not the issue. The order is in processing with enough inventory. The bottleneck may be operational (picking/packing) rather than a supply chain issue visible in the data.

## Must NOT Contain
- Claim that stock shortage is the cause
- Fabricated reasons not in the data

## Evaluation Checklist
- [ ] Order details correctly retrieved
- [ ] Inventory checked and found sufficient (35 available > 8 needed)
- [ ] No shipment confirmed
- [ ] Agent correctly notes stock is NOT the issue
- [ ] Agent acknowledges limitation (can't determine exact cause from available data)
- [ ] No hallucinated root cause
