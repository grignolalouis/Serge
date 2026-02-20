# Simple Supply Chain Model for Case Study

## Overview

We model a **small distribution company** that buys products from suppliers and sells to customers. This is the simplest complete supply chain that still demonstrates cross-source value.

```
┌──────────┐     ┌──────────────┐     ┌───────────┐     ┌───────────┐
│ Suppliers│────▶│  Warehouse   │────▶│ Transport │────▶│ Customers │
│  (3-5)   │     │   (1-2)      │     │  (fleet)  │     │  (50-100) │
└──────────┘     └──────────────┘     └───────────┘     └───────────┘
     │                  │                   │                 │
     ▼                  ▼                   ▼                 ▼
  Procurement       Inventory           Logistics          Sales
    Data              Data                Data              Data
```

---

## 1. Entities & Relationships

### Core Entities

```
┌─────────────────────────────────────────────────────────────────┐
│                        ENTITY MODEL                              │
├─────────────────────────────────────────────────────────────────┤
│                                                                  │
│  PRODUCT ─────────────┬─────────────── SUPPLIER                 │
│    │                  │                   │                      │
│    │ stored_in        │ supplied_by       │ delivers_to          │
│    ▼                  │                   ▼                      │
│  INVENTORY ◀──────────┘              PURCHASE_ORDER              │
│    │                                      │                      │
│    │ fulfills                             │ receives             │
│    ▼                                      ▼                      │
│  ORDER_LINE ◀──────── ORDER ◀──────── CUSTOMER                  │
│                         │                                        │
│                         │ shipped_via                            │
│                         ▼                                        │
│                     SHIPMENT ──────────▶ CARRIER                │
│                                                                  │
└─────────────────────────────────────────────────────────────────┘
```

---

## 2. Data Sources (Simulated Systems)

We simulate 4 distinct "systems" that the agent must query across:

| System | Real-World Equivalent | Data It Holds |
|--------|----------------------|---------------|
| **SRM** | Supplier Relationship Management | Suppliers, Purchase Orders, Supplier Performance |
| **WMS** | Warehouse Management System | Products, Inventory Levels, Stock Movements |
| **OMS** | Order Management System | Customers, Sales Orders, Order Lines |
| **TMS** | Transport Management System | Shipments, Carriers, Delivery Status |

### Why This Matters
The agent's value comes from **connecting these sources**. A human would need to check 4 different systems to answer: "Why is customer X's order delayed?"

---

## 3. Data Schema

### SRM (Supplier Data)

```sql
-- Suppliers
CREATE TABLE suppliers (
    id              UUID PRIMARY KEY,
    name            VARCHAR(100),
    country         VARCHAR(50),
    lead_time_days  INT,           -- Average delivery time
    reliability     DECIMAL(3,2),  -- 0.00-1.00 score
    contact_email   VARCHAR(100),
    created_at      TIMESTAMP
);

-- Purchase Orders (what we order from suppliers)
CREATE TABLE purchase_orders (
    id              UUID PRIMARY KEY,
    supplier_id     UUID REFERENCES suppliers(id),
    product_id      UUID,
    quantity        INT,
    unit_cost       DECIMAL(10,2),
    status          VARCHAR(20),   -- pending, confirmed, shipped, received, cancelled
    order_date      DATE,
    expected_date   DATE,
    received_date   DATE,
    created_at      TIMESTAMP
);
```

### WMS (Warehouse Data)

```sql
-- Products
CREATE TABLE products (
    id              UUID PRIMARY KEY,
    sku             VARCHAR(50) UNIQUE,
    name            VARCHAR(200),
    category        VARCHAR(50),
    unit_price      DECIMAL(10,2),
    weight_kg       DECIMAL(6,2),
    reorder_point   INT,           -- When to reorder
    reorder_qty     INT,           -- How much to reorder
    created_at      TIMESTAMP
);

-- Inventory (current stock per warehouse)
CREATE TABLE inventory (
    id              UUID PRIMARY KEY,
    product_id      UUID REFERENCES products(id),
    warehouse_id    VARCHAR(10),   -- 'WH-A', 'WH-B'
    quantity        INT,
    reserved        INT,           -- Reserved for orders
    available       INT GENERATED ALWAYS AS (quantity - reserved) STORED,
    last_updated    TIMESTAMP
);

-- Stock Movements (history)
CREATE TABLE stock_movements (
    id              UUID PRIMARY KEY,
    product_id      UUID REFERENCES products(id),
    warehouse_id    VARCHAR(10),
    movement_type   VARCHAR(20),   -- inbound, outbound, adjustment, transfer
    quantity        INT,           -- Positive or negative
    reference_id    UUID,          -- PO id or Order id
    reference_type  VARCHAR(20),   -- purchase_order, sales_order
    created_at      TIMESTAMP
);
```

### OMS (Order Data)

```sql
-- Customers
CREATE TABLE customers (
    id              UUID PRIMARY KEY,
    name            VARCHAR(100),
    company         VARCHAR(100),
    city            VARCHAR(50),
    country         VARCHAR(50),
    segment         VARCHAR(20),   -- retail, wholesale, enterprise
    created_at      TIMESTAMP
);

-- Sales Orders
CREATE TABLE orders (
    id              UUID PRIMARY KEY,
    customer_id     UUID REFERENCES customers(id),
    status          VARCHAR(20),   -- pending, confirmed, processing, shipped, delivered, cancelled
    order_date      DATE,
    required_date   DATE,          -- Customer wants it by
    shipped_date    DATE,
    delivered_date  DATE,
    total_amount    DECIMAL(12,2),
    priority        VARCHAR(10),   -- low, normal, high, urgent
    created_at      TIMESTAMP
);

-- Order Lines
CREATE TABLE order_lines (
    id              UUID PRIMARY KEY,
    order_id        UUID REFERENCES orders(id),
    product_id      UUID REFERENCES products(id),
    quantity        INT,
    unit_price      DECIMAL(10,2),
    status          VARCHAR(20),   -- pending, allocated, picked, packed, shipped
    created_at      TIMESTAMP
);
```

### TMS (Transport Data)

```sql
-- Carriers
CREATE TABLE carriers (
    id              UUID PRIMARY KEY,
    name            VARCHAR(100),
    type            VARCHAR(20),   -- ground, air, sea
    cost_per_kg     DECIMAL(6,2),
    avg_transit_days INT,
    on_time_rate    DECIMAL(3,2),  -- 0.00-1.00
    created_at      TIMESTAMP
);

-- Shipments
CREATE TABLE shipments (
    id              UUID PRIMARY KEY,
    order_id        UUID REFERENCES orders(id),
    carrier_id      UUID REFERENCES carriers(id),
    status          VARCHAR(20),   -- pending, picked_up, in_transit, out_for_delivery, delivered, exception
    tracking_number VARCHAR(50),
    origin          VARCHAR(50),
    destination     VARCHAR(100),
    weight_kg       DECIMAL(8,2),
    shipped_at      TIMESTAMP,
    estimated_arrival TIMESTAMP,
    delivered_at    TIMESTAMP,
    exception_reason VARCHAR(200), -- If status = exception
    created_at      TIMESTAMP
);
```

---

## 4. Sample Data Volume

| Entity | Count | Notes |
|--------|-------|-------|
| Suppliers | 5 | Different countries, varying reliability |
| Products | 30 | 3 categories, mix of fast/slow movers |
| Customers | 50 | Mix of segments |
| Orders | 500 | 3 months of history |
| Purchase Orders | 100 | From various suppliers |
| Shipments | 400 | Various statuses |
| Inventory Records | 60 | 30 products × 2 warehouses |

This is enough to demonstrate realistic queries without complexity.

---

## 5. Interesting Data Scenarios to Generate

Include these patterns in fake data for realistic agent queries:

### Scenario 1: Stockout Investigation
- Product X has 0 inventory
- Recent orders couldn't be fulfilled
- Purchase order from supplier was delayed

### Scenario 2: Late Delivery
- Customer order marked urgent
- Shipment shows "exception" status
- Carrier has low on-time rate

### Scenario 3: Supplier Comparison
- Same product available from 2 suppliers
- Different lead times, costs, reliability scores

### Scenario 4: Demand Spike
- Sudden increase in orders for product category
- Inventory depleting faster than reorder point predicted

### Scenario 5: Cross-Warehouse Transfer
- Product available in WH-B but not WH-A
- Customer is closer to WH-A

---

## 6. Data Relationships for Agent Queries

The agent must understand these **cross-source relationships**:

```
"Why is order #123 delayed?"
    │
    ├── OMS: Order status, required_date, customer info
    │
    ├── WMS: Were items in stock? When allocated?
    │
    └── TMS: Shipment status, carrier performance
            │
            └── If not shipped: Check SRM for pending PO

"What's the status of product SKU-456?"
    │
    ├── WMS: Current inventory levels per warehouse
    │
    ├── OMS: Pending orders requiring this product
    │
    ├── SRM: Incoming purchase orders
    │
    └── TMS: In-transit shipments containing this product
```

---

## 7. Data Generation Approach

### Option A: Script Generation (Recommended)
Write a Go or Python script that:
1. Generates consistent UUIDs
2. Creates realistic relationships
3. Includes the scenarios above
4. Outputs SQL or JSON seed files

### Option B: Use Faker Libraries
- Go: `github.com/brianvoe/gofakeit`
- Insert realistic names, dates, amounts

### Option C: LLM-Assisted Generation
- Use Claude to generate realistic order patterns
- Review and adjust for consistency
