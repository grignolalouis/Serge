# Spending by Supplier

## Use Case
UC-3: Supplier Comparison | Difficulty: medium

## Question
"How much are we spending with each supplier?"

## Expected Tool Calls
1. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-001")` — get POs for SUP-001
2. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-002")` — get POs for SUP-002
3. `srm_list_purchase_orders_by_supplier(supplier_id="SUP-003")` — get POs for SUP-003

Alternative: `srm_compare_suppliers()` then drill into POs.

## Ground Truth
**SUP-001 (Somchai Mango Farm):** 4 POs
- PO-001: 100 × 280 = 28,000 THB
- PO-004: 50 × 500 = 25,000 THB
- PO-006: 80 × 280 = 22,400 THB
- PO-008: 150 × 280 = 42,000 THB
- **Total: 117,400 THB**

**SUP-002 (Northern Fruits Co.):** 3 POs
- PO-002: 80 × 400 = 32,000 THB
- PO-005: 60 × 300 = 18,000 THB
- PO-009: 60 × 400 = 24,000 THB
- **Total: 74,000 THB**

**SUP-003 (Prachuap Harvest):** 3 POs
- PO-003: 120 × 200 = 24,000 THB
- PO-007: 100 × 200 = 20,000 THB
- PO-010: 80 × 300 = 24,000 THB
- **Total: 68,000 THB**

**Grand total: ~259,400 THB**

## Must NOT Contain
- Incorrect PO amounts
- Missing POs from any supplier

## Evaluation Checklist
- [ ] All 3 suppliers covered
- [ ] PO totals calculated correctly
- [ ] Per-supplier totals accurate
- [ ] Structured output (table or breakdown)
- [ ] No fabricated POs
