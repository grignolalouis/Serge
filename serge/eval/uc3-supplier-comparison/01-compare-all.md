# Compare All Suppliers

## Use Case
UC-3: Supplier Comparison | Difficulty: easy

## Question
"Compare all our mango suppliers."

## Expected Tool Calls
1. `srm_compare_suppliers()` — get side-by-side stats

## Ground Truth
| Supplier | ID | POs | Received | On-Time | On-Time Rate | Avg Lead |
|---|---|---|---|---|---|---|
| Somchai Mango Farm | SUP-001 | 4 | 3 | 3 | **100%** | **3 days** |
| Northern Fruits Co. | SUP-002 | 3 | 2 | 0 | **0%** | **8 days** |
| Prachuap Harvest | SUP-003 | 3 | 2 | 1 | **50%** | **6 days** |

- SUP-001: best performer — all deliveries on time or early, rating 4.5
- SUP-002: worst performer — both received POs late, rating 3.2
- SUP-003: mixed — one on time, one late, rating 2.8

## Must NOT Contain
- Missing any of the 3 suppliers
- Fabricated metrics
- Incorrect on-time rates

## Evaluation Checklist
- [ ] All 3 suppliers compared
- [ ] On-time rates correct for each
- [ ] Average lead times correct
- [ ] Clear ranking or recommendation
- [ ] Structured output (table or comparison format)
