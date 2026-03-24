# SERGE - Supply Chain Modelisation

## What is Being Simulated

SERGE V1 models a **Thai premium mango B2B distribution company** based in Bangkok. The company:

- **Sources** fresh mangoes from farms across Thailand
- **Stores** them in a central warehouse in Bangkok
- **Sells** to restaurants, hotels, supermarkets, and food halls
- **Ships** via ground and refrigerated cold-chain carriers

This is a **single-product-category, single-warehouse, domestic B2B** supply chain -- intentionally simple to validate the AI agent before scaling complexity.

## V1 Supply Chain Flow

```
[Thai Mango Farms]         [Bangkok Warehouse]         [B2B Customers]
  3 suppliers       --->     WH-BKK (Bang Na)    --->   50 customers
  Chanthaburi                5 product SKUs              Restaurants
  Chiang Mai                 Inventory tracking           Hotels
  Prachuap Khiri Khan        Reorder points               Supermarkets
                                                          Food halls
        |                         |                          |
   Purchase Orders           Stock Movements           Sales Orders
   (SRM system)              (WMS system)              (OMS system)
                                  |
                             [2 Carriers]
                          Kerry Express (ground)
                          SCG Cold Chain (refrigerated)
                                  |
                             Shipments + Tracking
                             (TMS system)
```

## V1 Data Model

### Products (5 SKUs)
| SKU | Name | Category | Price (THB) | Weight |
|-----|------|----------|-------------|--------|
| MANGO-NDM-5 | Nam Doc Mai Mango Box 5kg | Premium | 450 | 5kg |
| MANGO-ORG-10 | Ok Rong Mango Box 10kg | Standard | 650 | 10kg |
| MANGO-KEW-5 | Keaw Mango Box 5kg | Standard | 350 | 5kg |
| MANGO-MHC-10 | Mahachanok Mango Box 10kg | Premium | 800 | 10kg |
| MANGO-MIX-5 | Mixed Mango Assortment 5kg | Assorted | 500 | 5kg |

### Suppliers (3)
| Supplier | Region | Lead Time | Rating |
|----------|--------|-----------|--------|
| Somchai Mango Farm | Chanthaburi | 3 days | 4.5/5 |
| Northern Fruits Co. | Chiang Mai | 5 days | 3.2/5 |
| Prachuap Harvest | Prachuap Khiri Khan | 4 days | 2.8/5 |

### Customers (50)
- Segments: **retail** (supermarkets, food halls) and **wholesale** (restaurants, hotels)
- Region: primarily Bangkok
- Examples: Siam Paragon Food Hall, Blue Elephant Restaurant, Mandarin Oriental Bangkok, Tops Supermarket

### Carriers (2)
| Carrier | Type | Cost/kg | Avg Transit |
|---------|------|---------|-------------|
| Kerry Express | Ground | 15 THB | 2 days |
| SCG Cold Chain | Refrigerated | 25 THB | 1 day |

## V1 Scale
- 500 orders, 400 shipments, ~1000 tracking events
- Static seed data (read-only, loaded at startup)
- Single warehouse, domestic only

## Why This Model for V1

The V1 modelisation is intentionally constrained to:

1. **Validate the agent architecture** -- prove that an LLM can reason across 4 systems (SRM, WMS, OMS, TMS) via MCP tools
2. **Build evaluation benchmarks** -- with a small, known dataset, we can define ground truth for LLM-as-judge evaluation
3. **Demonstrate cross-system reasoning** -- e.g., "why is order X delayed?" requires checking OMS (order status) -> TMS (no shipment) -> WMS (insufficient stock) -> SRM (supplier performance)

## Roadmap: Simple to Real B2B Simulation

### V2 - Dynamic Operations
- **Write operations**: create purchase orders, update inventory, advance order status
- **Time simulation**: clock that advances, triggering events (deliveries arrive, stock depletes)
- **Multi-warehouse**: add Chiang Mai distribution center, regional routing
- **More products**: expand beyond mangoes to tropical fruits (durian, longan, rambutan)

### V3 - Realistic B2B Complexity
- **Demand forecasting**: seasonal mango patterns (peak: Mar-Jun), historical trends
- **Contract management**: supplier agreements, volume discounts, SLAs
- **Payment terms**: Net-30, credit limits, invoice tracking
- **Returns & quality**: rejection rates, quality grades, cold chain compliance
- **Multi-carrier routing**: cost vs speed optimization, carrier capacity constraints

### V4 - Full B2B Simulation Company
- **International trade**: export to Japan, Singapore, Middle East (customs, duties, phytosanitary certificates)
- **Real-time events**: weather disruptions, supplier outages, demand spikes
- **Financial integration**: cost accounting, margin analysis, cash flow impact
- **Compliance**: food safety (GMP, HACCP), traceability requirements
- **Multi-tenant**: simulate multiple business units or subsidiaries
- **API integration**: connect to real ERP/WMS systems instead of seed data

### Evaluation Progression
| Version | Complexity | Evaluation Focus |
|---------|-----------|-----------------|
| V1 | Static, read-only, small dataset | Tool selection accuracy, cross-system reasoning, factual correctness |
| V2 | Dynamic, write-enabled | Action planning, multi-step workflows, state management |
| V3 | Realistic constraints | Decision quality, trade-off reasoning, anomaly detection |
| V4 | Full simulation | End-to-end business value, proactive insights, ROI measurement |
