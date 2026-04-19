# Product Lookup

## Metadata
- **Level:** L1
- **Use Case:** UC-4 (operational)
- **Systems:** WMS
- **Expected tools:** 1
- **Realistic user:** Anyone needing basic product info

## Question
"Can you tell me about product PROD-001?"

## Required Facts

**Critical:**
- Name: Nam Doc Mai Mango Box 5kg
- SKU: MANGO-NDM-5
- Category: premium
- Unit price: 450 THB

**Important:**
- Weight: 5 kg
- Shelf life: 7 days
- Storage: refrigerated
- Reorder point: 20

**Nice to have:**
- Any commentary on perishability or premium positioning

## Forbidden Content
- Wrong price, SKU, or product name
- Fabricated supplier info (answer should not invent which suppliers sell it unless another tool is called)

## Flexibility Notes
- Agent may call `wms_get_product` or `wms_get_product_by_sku` — both valid
- Format can be prose, bullet list, or table

## Ground Truth

| Key | Value |
|-----|-------|
| id | PROD-001 |
| sku | MANGO-NDM-5 |
| name | Nam Doc Mai Mango Box 5kg |
| category | premium |
| unit_price | 450 |
| weight_kg | 5 |
| reorder_point | 20 |
| shelf_life_days | 7 |
| storage_type | refrigerated |
