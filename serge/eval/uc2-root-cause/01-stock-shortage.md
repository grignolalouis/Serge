# Stock Shortage — Order Blocked by Low Inventory

## Use Case
UC-2: Root Cause Analysis | Difficulty: hard

## Question
"Why hasn't order ORD-007 shipped yet?"

## Expected Tool Calls
1. `oms_get_order(order_id="ORD-007")` — see order is stuck in processing
2. `tms_track_order(order_id="ORD-007")` — confirm no shipment exists
3. `wms_check_inventory(product_id="PROD-001")` — discover only 8 units available, need 50
4. `srm_list_purchase_orders_by_status(status="shipped")` — find PO-008 incoming with PROD-001

## Ground Truth
- ORD-007: CUST-001 (Siam Paragon), status **processing**, priority **urgent**
- Ordered 2026-02-05, required by 2026-02-12 — **overdue**
- Needs: 50x PROD-001 (Nam Doc Mai 5kg) @ 450 THB = 22,500 THB
- No shipment created
- WMS: PROD-001 has **8 units available**, 0 reserved, reorder point 20 — **critically low**
- SRM: PO-008 from SUP-001 (Somchai Mango Farm), status **shipped**, 150 units PROD-001, ETA **2026-02-18**
- **Root cause**: inventory shortage (8 available, 50 needed). Resupply PO-008 arrives Feb 18, 6 days after required date.
- The order **cannot be fulfilled** until PO-008 is received.

## Must NOT Contain
- Incorrect root cause (e.g., carrier issue, cancelled order)
- Missing causal link between inventory and PO
- Claim that the order can be fulfilled from current stock

## Evaluation Checklist
- [ ] All 4 systems queried (OMS → TMS → WMS → SRM)
- [ ] Root cause correctly identified as inventory shortage
- [ ] Specific numbers: 8 available vs 50 needed
- [ ] PO-008 referenced as incoming resupply
- [ ] Gap noted: PO arrives Feb 18, order required Feb 12
- [ ] Causal chain is complete and logical
