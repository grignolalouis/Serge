# SERGE — Data Model

```
┌─────────────────────────────────────────────────────────────────────────────────────────────┐
│                                    PROCUREMENT (SRM)                                        │
│                                                                                             │
│  ┌──────────────────────┐         ┌──────────────────────────┐                              │
│  │      Supplier        │         │    SupplierProduct       │                              │
│  │──────────────────────│         │──────────────────────────│                              │
│  │ id           SUP-001 │◄───┐    │ id           SP-001      │                              │
│  │ name                 │    ├────│ supplier_id  SUP-001 (FK)│                              │
│  │ region               │    │    │ product_id   PROD-001(FK)│───────────────┐              │
│  │ contact              │    │    │ name                     │               │              │
│  │ lead_time_days       │    │    │ unit_cost                │               │              │
│  │ rating          1-5  │    │    └──────────────────────────┘               │              │
│  │ product_categories[] │    │                 ▲                             │              │
│  │ payment_terms        │    │                 │ supplier_product_id         │              │
│  └──────────────────────┘    │                 │                             │              │
│           ▲                  │    ┌────────────┴─────────────┐               │              │
│           │ supplier_id      │    │  SupplierOrderLine       │               │              │
│           │                  │    │──────────────────────────│               │              │
│  ┌────────┴─────────────┐    │    │ supplier_product_id (FK) │               │              │
│  │   SupplierOrder      │    │    │ quantity                 │               │              │
│  │──────────────────────│    │    └──────────────────────────┘               │              │
│  │ id           PO-001  │────┘                 ▲                             │              │
│  │ supplier_id  (FK)    │                      │ embedded                    │              │
│  │ status               │──────────────────────┘                             │              │
│  │ order_date           │                                                    │              │
│  │ expected_delivery    │                                                    │              │
│  │ actual_delivery      │                                                    │              │
│  │ lines[]              │                                                    │              │
│  └──────────┬───────────┘                                                    │              │
│             │                                                                │              │
└─────────────│────────────────────────────────────────────────────────────────│──────────────┘
              │ reference_id (inbound)                                         │
              ▼                                                                ▼
┌─────────────────────────────────────────────────────────────────────────────────────────────┐
│                                     INVENTORY (WMS)                                         │
│                                                                                             │
│  ┌──────────────────────┐    ┌───────────────────────────┐    ┌──────────────────────┐      │
│  │     Warehouse        │    │     InventoryRecord       │    │      Product         │      │
│  │──────────────────────│    │───────────────────────────│    │──────────────────────│      │
│  │ id          WH-BKK  │◄───│ warehouse_id    (FK)      │    │ id         PROD-001  │◄─────┤
│  │ name                 │    │ product_id      (FK)      │───►│ sku                  │      │
│  │ location             │    │ quantity                  │    │ name                 │      │
│  └──────────────────────┘    │ reserved                  │    │ category             │      │
│           ▲                  │ last_updated              │    │ unit_price           │      │
│           │ warehouse_id     └───────────────────────────┘    │ weight_kg            │      │
│           │                                                   │ reorder_point        │      │
│  ┌────────┴─────────────┐                                     │ shelf_life_days      │      │
│  │   StockMovement      │                                     │ storage_type         │      │
│  │──────────────────────│                                     └──────────────────────┘      │
│  │ id           SM-001  │                                              ▲                    │
│  │ product_id   (FK)    │──────────────────────────────────────────────┘                    │
│  │ warehouse_id (FK)    │                                              ▲                    │
│  │ type  inbound/outb.  │                                              │                    │
│  │ quantity             │                                              │ product_id         │
│  │ reference_id  (poly) │─── PO-xxx (inbound) or ORD-xxx (outbound)   │                    │
│  │ reference_type       │                                              │                    │
│  │ created_at           │                                              │                    │
│  └──────────┬───────────┘                                              │                    │
│             │                                                          │                    │
└─────────────│──────────────────────────────────────────────────────────│────────────────────┘
              │ reference_id (outbound)                                  │
              ▼                                                          │
┌─────────────────────────────────────────────────────────────────────────────────────────────┐
│                                       SALES (OMS)                                           │
│                                                                                             │
│  ┌──────────────────────┐         ┌──────────────────────────┐                              │
│  │      Customer        │         │  CustomerOrderLine       │                              │
│  │──────────────────────│         │──────────────────────────│                              │
│  │ id         CUST-001  │◄───┐    │ product_id   PROD-001(FK)│──────────────────────────────┘
│  │ name                 │    │    │ quantity                 │
│  │ company              │    │    └──────────────────────────┘
│  │ segment  retail/whol │    │                 ▲
│  │ region               │    │                 │ embedded
│  │ address              │    │                 │
│  │ delivery_zone        │    │    ┌────────────┴─────────────┐
│  └──────────────────────┘    │    │    CustomerOrder         │
│                              │    │──────────────────────────│
│                              └────│ customer_id      (FK)    │
│                                   │ id           ORD-001     │
│                                   │ status                   │
│                                   │ order_date               │
│                                   │ required_date            │
│                                   │ shipped_date             │
│                                   │ delivered_date           │
│                                   │ priority  norm/high/urg  │
│                                   │ notes                    │
│                                   │ lines[]                  │
│                                   └────────────┬─────────────┘
│                                                │
└────────────────────────────────────────────────│────────────────────────────────────────────┘
                                                 │ customer_order_id
                                                 ▼
┌─────────────────────────────────────────────────────────────────────────────────────────────┐
│                                      SHIPPING (TMS)                                         │
│                                                                                             │
│  ┌──────────────────────┐         ┌──────────────────────────┐                              │
│  │      Carrier         │         │      Shipment            │                              │
│  │──────────────────────│         │──────────────────────────│                              │
│  │ id         CAR-001   │◄────────│ carrier_id       (FK)    │                              │
│  │ name                 │         │ id           SHP-001     │                              │
│  │ type  ground/refrig  │         │ customer_order_id (FK)   │                              │
│  │ cost_per_kg          │         │ status                   │                              │
│  │ avg_transit_days     │         │ tracking_number          │                              │
│  └──────────────────────┘         │ origin                   │                              │
│                                   │ destination              │                              │
│                                   │ weight_kg                │                              │
│                                   │ shipped_at               │                              │
│                                   │ estimated_arrival        │                              │
│                                   │ delivered_at             │                              │
│                                   │ exception_reason         │                              │
│                                   └────────────┬─────────────┘                              │
│                                                │                                            │
│                                                │ shipment_id                                │
│                                                ▼                                            │
│                                   ┌──────────────────────────┐                              │
│                                   │    TrackingEvent         │                              │
│                                   │──────────────────────────│                              │
│                                   │ shipment_id      (FK)    │                              │
│                                   │ timestamp                │                              │
│                                   │ location                 │                              │
│                                   │ status                   │                              │
│                                   │ description              │                              │
│                                   └──────────────────────────┘                              │
│                                                                                             │
└─────────────────────────────────────────────────────────────────────────────────────────────┘
```

## Cross-System Foreign Keys

```
PROCUREMENT (SRM)                INVENTORY (WMS)                SALES (OMS)               SHIPPING (TMS)

Supplier ◄──── SupplierProduct ────► Product ◄──── CustomerOrderLine ──── CustomerOrder ────► Shipment ────► Carrier
  ▲                  ▲                  ▲                                       ▲                │
  │                  │                  │                                       │                │
  └── SupplierOrder  │           InventoryRecord                               │          TrackingEvent
        │            │                  │                                       │
        │       SupplierOrderLine    StockMovement                              │
        │                               │                                      │
        └──── reference_id (inbound) ───┘──── reference_id (outbound) ─────────┘
```

## Relationships Summary

| From Entity        | Field                | To Entity       | Cardinality | Notes                           |
|--------------------|----------------------|-----------------|-------------|---------------------------------|
| SupplierProduct    | supplier_id          | Supplier        | N:1         | Many products per supplier      |
| SupplierProduct    | product_id           | Product         | N:1         | Many suppliers per product      |
| SupplierOrder      | supplier_id          | Supplier        | N:1         | Many POs per supplier           |
| SupplierOrderLine  | supplier_product_id  | SupplierProduct | N:1         | Line references supplier's item |
| InventoryRecord    | product_id           | Product         | 1:1         | One record per product/warehouse|
| InventoryRecord    | warehouse_id         | Warehouse       | N:1         | Many records per warehouse      |
| StockMovement      | product_id           | Product         | N:1         | Movement tracks a product       |
| StockMovement      | warehouse_id         | Warehouse       | N:1         | Movement in/out of warehouse    |
| StockMovement      | reference_id (poly)  | SupplierOrder   | N:1         | Inbound: from PO               |
| StockMovement      | reference_id (poly)  | CustomerOrder   | N:1         | Outbound: for customer order    |
| CustomerOrder      | customer_id          | Customer        | N:1         | Many orders per customer        |
| CustomerOrderLine  | product_id           | Product         | N:1         | Line references our product     |
| Shipment           | customer_order_id    | CustomerOrder   | N:1         | Many shipments per order (split)|
| Shipment           | carrier_id           | Carrier         | N:1         | One carrier per shipment        |
| TrackingEvent      | shipment_id          | Shipment        | N:1         | Many events per shipment        |

## Enums

| Entity          | Field          | Values                                       |
|-----------------|----------------|----------------------------------------------|
| SupplierOrder   | status         | pending, confirmed, shipped, received, cancelled |
| CustomerOrder   | status         | pending, processing, shipped, delivered, cancelled |
| CustomerOrder   | priority       | normal, high, urgent                         |
| Shipment        | status         | pending, in_transit, delivered, exception     |
| StockMovement   | type           | inbound, outbound                            |
| StockMovement   | reference_type | supplier_order, customer_order               |
| Product         | storage_type   | refrigerated, ambient                        |
| Carrier         | type           | ground, refrigerated                         |
| Supplier        | payment_terms  | net_15, net_30, cod                          |
| Customer        | segment        | retail, wholesale                            |

## Data Volumes (Seed)

| Entity           | Count | ID Pattern |
|------------------|------:|------------|
| Supplier         |     5 | SUP-001    |
| SupplierProduct  |    15 | SP-001     |
| Product          |     8 | PROD-001   |
| Warehouse        |     1 | WH-BKK    |
| Customer         |    12 | CUST-001   |
| Carrier          |     3 | CAR-001    |
| SupplierOrder    |    60 | PO-001     |
| CustomerOrder    |   500 | ORD-001    |
| Shipment         |   493 | SHP-001    |
| TrackingEvent    | 1,431 | —          |
| StockMovement    |   965 | SM-001     |
| InventoryRecord  |     8 | —          |

## Business Rules

- **Price model**: purchase cost on `SupplierProduct.unit_cost`, selling price on `Product.unit_price`. Order lines carry quantity only.
- **Refrigeration rule**: `Product.storage_type = refrigerated` **must** ship via `Carrier.type = refrigerated`.
- **Stock availability**: `InventoryRecord.quantity - reserved`. Alert when available < `Product.reorder_point`.
- **Split shipments**: one `CustomerOrder` can have multiple `Shipment` records (e.g., fresh via cold chain, dried via ground).
- **Polymorphic reference**: `StockMovement.reference_id` points to either a `SupplierOrder` (inbound) or `CustomerOrder` (outbound), disambiguated by `reference_type`.
