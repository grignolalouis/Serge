# Product Demand View

## Metadata
- **Level:** L3
- **Use Case:** UC-4
- **Systems:** OMS, WMS
- **Expected tools:** 2
- **Realistic user:** Ops manager checking demand vs supply

## Question
"Which orders are requesting PROD-001 and do we have enough stock to cover them?"

## Required Facts

**Critical:**
- Current PROD-001 inventory: 8 available
- Multiple scenario orders need PROD-001: ORD-007 (50 units), ORD-009 (5), others from ORD-004, ORD-015, etc. (already delivered)
- Pending/processing demand exceeds current stock

**Important:**
- Total unfulfilled demand from processing/pending/shipped orders
- ORD-007 alone (50 units) exceeds current stock

**Nice to have:**
- Mention of PO-008 (150 units incoming)
- Severity assessment

## Forbidden Content
- Wrong available stock
- Orders not containing PROD-001

## Flexibility Notes
- Agent may list many orders (generated data contains lots of historical orders with PROD-001). The judge focuses on correct arithmetic — the agent should distinguish already-delivered orders from active demand.
- Counting every single historical order is not required; the key is active demand vs stock.

## Ground Truth

| Key | Value |
|-----|-------|
| PROD-001 available | 8 |
| PROD-001 reorder point | 20 |
| ORD-007 needs | 50 units (processing) |
| ORD-009 needs | 5 units (pending) |
| Conclusion | Stock cannot cover active demand |
