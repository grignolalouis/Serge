# Product With Current Inventory

## Metadata
- **Level:** L3
- **Use Case:** UC-4
- **Systems:** WMS
- **Expected tools:** 1–2
- **Realistic user:** Warehouse manager checking a specific item

## Question
"Full inventory status for PROD-004 — what's on hand and do we need to reorder?"

## Required Facts

**Critical:**
- Quantity on hand: 22
- Reserved: 3
- Available: 19
- Reorder point: 10
- **No reorder needed** (19 > 10)

**Important:**
- Product name: Mahachanok Mango Box 10kg
- Category: premium
- Shelf life: 5 days (shortest of all fresh mangoes — worth noting)

**Nice to have:**
- Comment on short shelf life implications
- Storage type (refrigerated)

## Forbidden Content
- Claim reorder is needed (it isn't)
- Wrong inventory figures

## Flexibility Notes
- `wms_check_inventory` returns everything in one call
- Agent may also fetch recent stock movements for trend context (bonus, not required)

## Ground Truth

| Key | Value |
|-----|-------|
| product_id | PROD-004 |
| quantity | 22 |
| reserved | 3 |
| available | 19 |
| reorder_point | 10 |
| needs_reorder | false |
| shelf_life_days | 5 |
