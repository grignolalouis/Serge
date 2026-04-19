# Low Stock Alert

## Metadata
- **Level:** L2
- **Use Case:** UC-4
- **Systems:** WMS
- **Expected tools:** 1
- **Realistic user:** Warehouse manager doing morning review

## Question
"What products are below their reorder point right now?"

## Required Facts

**Critical:**
- PROD-001 (Nam Doc Mai) — 8 available vs reorder point 20
- PROD-005 (Mixed Mango Assortment) — 12 available vs reorder point 15
- PROD-008 (Sticky Rice Kit) — 6 available vs reorder point 15

**Important:**
- Each alert includes the gap (how many units short)

**Nice to have:**
- Prioritization by severity
- Suggestion to place a supplier order

## Forbidden Content
- Products that are NOT below reorder point (PROD-002, PROD-003, PROD-004, PROD-006, PROD-007)
- Wrong available quantities

## Flexibility Notes
- The judge accepts either product IDs or names
- Adding a "severity" ranking is bonus, not required

## Ground Truth

| Product | Available | Reorder Point | Needs Reorder |
|---------|-----------|---------------|---------------|
| PROD-001 | 8 | 20 | yes (critical) |
| PROD-005 | 12 | 15 | yes |
| PROD-008 | 6 | 15 | yes |
