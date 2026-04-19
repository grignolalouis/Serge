# Supplier Side-by-Side Comparison

## Metadata
- **Level:** L3
- **Use Case:** UC-3
- **Systems:** SRM
- **Expected tools:** 1–3
- **Realistic user:** Procurement manager evaluating sourcing

## Question
"Compare our suppliers — who is the most reliable?"

## Required Facts

**Critical:**
- SUP-001 (Somchai Mango Farm): high on-time rate (100% for received scenario POs), short avg lead
- SUP-002 (Northern Fruits Co.): low on-time rate (0% — both received POs late), long avg lead
- SUP-003 (Prachuap Harvest): mixed (1 on-time, 1 late among received)
- SUP-004 (Rayong Tropical): solid (on-time or early)
- SUP-005 (Isaan Organic): some delay on at least one received PO

**Important:**
- A ranking or verdict (best → worst)
- Approximate lead times

**Nice to have:**
- Rating field values
- Cost considerations

## Forbidden Content
- Fabricated on-time rates
- Missing suppliers (only 5 exist: SUP-001 through SUP-005)

## Flexibility Notes
- The `srm_compare_suppliers` tool returns everything needed in one call — efficient path
- Agent may also drill into individual suppliers; both valid
- Exact percentages may vary slightly due to generated PO data (IDs ≥ 100); the judge focuses on the relative ordering, not precise decimals

## Ground Truth (relative, from scenario data)

| Supplier | Scenario On-Time | Avg Lead (scenario) | Rating |
|----------|------------------|---------------------|--------|
| SUP-001 | 100% (3/3) | ~3 days | 4.5 |
| SUP-002 | 0% (0/2) | ~8 days | 3.2 |
| SUP-003 | 50% (1/2) | ~6 days | 2.8 |
| SUP-004 | high (scenario POs on-time or early) | ~3 days | 4.0 |
| SUP-005 | mixed | ~7 days | 3.5 |

Best reliability: **SUP-001** (Somchai Mango Farm).
