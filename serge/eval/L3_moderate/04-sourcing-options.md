# Sourcing Options for a Product

## Metadata
- **Level:** L3
- **Use Case:** UC-3
- **Systems:** SRM, WMS
- **Expected tools:** 2–3
- **Realistic user:** Procurement looking for alternatives

## Question
"Who can supply us Nam Doc Mai mangoes and at what price?"

## Required Facts

**Critical:**
- PROD-001 (Nam Doc Mai) has 2 suppliers: SUP-001 and SUP-004
- SUP-001 cost: 280 THB (SP-001)
- SUP-004 cost: 270 THB (SP-009)

**Important:**
- SUP-001 (Somchai Mango Farm, Chanthaburi) rating 4.5, lead time 3 days
- SUP-004 (Rayong Tropical) rating 4.0, lead time 3 days

**Nice to have:**
- Cost comparison and verdict (SUP-004 slightly cheaper, SUP-001 slightly higher-rated)
- Our selling price (450 THB) and resulting margin

## Forbidden Content
- Fabricated suppliers
- Wrong unit costs

## Flexibility Notes
- `srm_list_supplier_products_by_product` with product_id=PROD-001 returns both options directly — cleanest path
- Agent may also fetch full supplier profiles for context

## Ground Truth

| Supplier Product | Supplier | Unit Cost |
|------------------|----------|-----------|
| SP-001 | SUP-001 Somchai Mango Farm | 280 THB |
| SP-009 | SUP-004 Rayong Tropical | 270 THB |
