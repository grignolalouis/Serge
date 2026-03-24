# Late Delivery History

## Use Case
UC-3: Supplier Comparison | Difficulty: medium

## Question
"Have any of our suppliers caused problems with late deliveries?"

## Expected Tool Calls
1. `srm_compare_suppliers()` — overview
2. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-002")` — drill into worst performer
3. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-003")` — check mid performer

## Ground Truth
**SUP-002 (Northern Fruits Co.) — worst offender:**
- PO-002: expected Jan 10, received Jan 13 → **3 days late**
- PO-005: expected Jan 20, received Jan 24 → **4 days late**
- 0% on-time rate, both deliveries late
- Pattern: consistently late, getting worse (3 → 4 days)

**SUP-003 (Prachuap Harvest) — occasional issues:**
- PO-003: expected Jan 12, received Jan 12 → on time
- PO-007: expected Jan 29, received Feb 2 → **4 days late**
- 50% on-time rate, inconsistent

**SUP-001 (Somchai Mango Farm) — no issues:**
- PO-001: on time, PO-004: 1 day early, PO-006: on time
- 100% on-time rate

## Must NOT Contain
- Claim SUP-001 has late deliveries
- Missing any late PO
- Wrong delay durations

## Evaluation Checklist
- [ ] SUP-002 identified as main problem (0% on-time)
- [ ] Specific late POs cited with dates and delay duration
- [ ] SUP-003 noted as occasional issue
- [ ] SUP-001 confirmed as reliable
- [ ] Pattern/trend noted for SUP-002
- [ ] No fabricated data
