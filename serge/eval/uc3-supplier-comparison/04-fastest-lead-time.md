# Fastest Lead Time

## Use Case
UC-3: Supplier Comparison | Difficulty: easy

## Question
"Which supplier delivers fastest on average?"

## Expected Tool Calls
1. `srm_compare_suppliers()` — get avg lead time stats

## Ground Truth
Based on received POs (actual order date → actual delivery):

**SUP-001 (Somchai Mango Farm): ~3 days average**
- PO-001: Jan 2 → Jan 5 = 3 days
- PO-004: Jan 12 → Jan 14 = 2 days
- PO-006: Jan 22 → Jan 25 = 3 days
- Average: (3+2+3)/3 = **2.7 days** ≈ 3 days

**SUP-003 (Prachuap Harvest): ~6 days average**
- PO-003: Jan 8 → Jan 12 = 4 days
- PO-007: Jan 25 → Feb 2 = 8 days
- Average: (4+8)/2 = **6 days**

**SUP-002 (Northern Fruits Co.): ~8 days average**
- PO-002: Jan 5 → Jan 13 = 8 days
- PO-005: Jan 15 → Jan 24 = 9 days
- Average: (8+9)/2 = **8.5 days** ≈ 8 days

**Fastest: SUP-001** at ~3 days, nearly 3x faster than SUP-002.

## Must NOT Contain
- Wrong supplier ranked as fastest
- Incorrect lead time calculations

## Evaluation Checklist
- [ ] SUP-001 correctly identified as fastest
- [ ] Lead times approximately correct for each
- [ ] Comparison with declared lead times noted (SUP-001 declared 3 days, actual ~3)
- [ ] Structured ranking
