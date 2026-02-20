# Use Cases for Supply Chain Agent

## Overview

These use cases demonstrate the **cross-source value** of the agent. Each requires data from multiple systems that a human would need to query separately.

---

## Use Case 1: Order Delay Investigation

### Scenario
A customer calls asking why their order hasn't arrived.

### Without Agent
1. Check OMS for order status
2. Check TMS for shipment status
3. If not shipped, check WMS for inventory
4. If inventory short, check SRM for pending POs
5. Manually piece together the story

**Time: 15-30 minutes**

### With Agent

```
User: "Why is order #ORD-2024-0456 delayed?"

Agent: Based on my analysis across systems:

**Order Status** (OMS):
- Order #ORD-2024-0456 placed Jan 10, required by Jan 15
- Status: Processing (not yet shipped)
- Contains: 50x SKU-WIDGET-A, 20x SKU-GADGET-B

**Inventory Issue** (WMS):
- SKU-WIDGET-A: Only 12 units available (need 50)
- SKU-GADGET-B: 45 units available ✓

**Replenishment Status** (SRM):
- PO-2024-089 for 100x SKU-WIDGET-A from Acme Supplies
- Status: In transit, expected arrival: Jan 16

**Root Cause**: Order is waiting on inventory replenishment for SKU-WIDGET-A.

**Expected Resolution**: Order can ship Jan 17 (after PO receipt + processing).

**Recommendation**: Contact customer to offer:
1. Partial shipment of SKU-GADGET-B now
2. Expedited shipping once complete
```

**Time: 30 seconds**

---

## Use Case 2: Product Availability Check

### Scenario
Sales team needs to know if they can promise delivery of 200 units to a large customer.

### Query

```
User: "Can we fulfill 200 units of SKU-MOTOR-X by next Friday?"

Agent: Let me check availability across all sources:

**Current Inventory** (WMS):
- Warehouse A: 75 units available
- Warehouse B: 45 units available
- Total available: 120 units

**Committed Stock** (OMS):
- 30 units reserved for pending orders
- Net available: 90 units

**Incoming Supply** (SRM):
- PO-2024-102: 150 units arriving Jan 18 (before next Friday)
- Supplier reliability: 94%

**Analysis**:
- Current: 90 units available
- After PO: 240 units available
- Requested: 200 units
- Buffer: 40 units

**Answer**: YES, we can fulfill 200 units by next Friday.

**Risk**: Depends on PO-2024-102 arriving on time. Supplier (TechParts Inc)
has 94% on-time rate. Recommend confirming shipment status mid-week.
```

---

## Use Case 3: Supplier Performance Comparison

### Scenario
Procurement needs to decide which supplier to use for a new bulk order.

### Query

```
User: "Compare our top 3 suppliers for electronic components"

Agent: Here's a comparison based on the last 6 months of data:

| Metric | Acme Electronics | TechParts Inc | GlobalSupply Co |
|--------|-----------------|---------------|-----------------|
| **Orders Placed** | 45 | 38 | 52 |
| **On-Time Delivery** | 91% | 96% | 78% |
| **Average Lead Time** | 12 days | 8 days | 15 days |
| **Quality Issues** | 3 returns | 1 return | 7 returns |
| **Avg Unit Cost** | $24.50 | $26.00 | $22.00 |
| **Total Spend** | $142,000 | $98,000 | $178,000 |

**Cost vs Reliability Trade-off**:
- GlobalSupply Co is cheapest but has reliability issues
- TechParts Inc has best delivery performance (+6% premium)
- Acme Electronics is balanced option

**Recommendation**: For time-sensitive orders, use TechParts Inc.
For bulk orders with flexibility, consider GlobalSupply Co with buffer stock.
```

---

## Use Case 4: Shipment Exception Handling

### Scenario
Logistics coordinator needs to handle delivery exceptions for the day.

### Query

```
User: "Show me all shipment exceptions and their impact"

Agent: Found 4 shipments with exceptions:

**1. SHP-2024-1892** ⚠️ HIGH PRIORITY
- Order: #ORD-2024-0501 (Customer: MegaCorp)
- Issue: Customs hold in Frankfurt
- Carrier: DHL Express
- Customer impact: Order marked URGENT, 2 days overdue
- Action needed: Contact DHL for clearance status

**2. SHP-2024-1901** ⚠️ MEDIUM
- Order: #ORD-2024-0498 (Customer: TechRetail)
- Issue: Address verification failed
- Carrier: FedEx
- Customer impact: Normal priority, within SLA
- Action needed: Verify address with customer

**3. SHP-2024-1905** ⚠️ MEDIUM
- Order: #ORD-2024-0503 (Customer: SmallBiz LLC)
- Issue: Package damaged in transit
- Carrier: UPS
- Customer impact: Partial damage, may need replacement
- Inventory check: 15 units available for re-ship

**4. SHP-2024-1910** ⚠️ LOW
- Order: #ORD-2024-0510 (Customer: LocalStore)
- Issue: Delivery attempted, recipient unavailable
- Carrier: Local Courier
- Action needed: Schedule re-delivery

**Summary**: 4 exceptions, 1 requires immediate attention (MegaCorp customs hold).
```

---

## Use Case 5: Low Stock Alert Analysis

### Scenario
Morning check on inventory health.

### Query

```
User: "What products are at risk of stockout this week?"

Agent: Analyzing inventory vs demand patterns:

**Critical - Will stockout in 0-2 days:**

| SKU | Current Stock | Daily Demand | Days Left | Incoming PO |
|-----|--------------|--------------|-----------|-------------|
| SKU-CHIP-22 | 45 | 28 | 1.6 days | PO-2024-115 due Jan 20 (5 days) |
| SKU-CABLE-USB | 12 | 15 | 0.8 days | None scheduled |

**Warning - Will stockout in 3-5 days:**

| SKU | Current Stock | Daily Demand | Days Left | Incoming PO |
|-----|--------------|--------------|-----------|-------------|
| SKU-SENSOR-T1 | 89 | 22 | 4 days | PO-2024-118 due Jan 18 (3 days) ✓ |
| SKU-MOUNT-BR | 156 | 35 | 4.5 days | None scheduled |

**Immediate Actions Required:**

1. **SKU-CHIP-22**: Gap between stockout (1.6 days) and PO arrival (5 days).
   → Contact supplier for expedite OR find alternate supplier

2. **SKU-CABLE-USB**: No PO scheduled.
   → Create emergency purchase order immediately
   → Suggested supplier: TechParts Inc (fastest lead time: 3 days)

3. **SKU-MOUNT-BR**: Trending to stockout with no replenishment.
   → Review if demand spike is temporary or needs reorder point adjustment
```

---

## Use Case 6: Customer Order History

### Scenario
Account manager preparing for a customer meeting.

### Query

```
User: "Give me a summary of MegaCorp's account for my meeting tomorrow"

Agent: **MegaCorp Account Summary**

**Customer Profile** (OMS):
- Segment: Enterprise
- Account since: March 2022
- Primary contact: John Smith, j.smith@megacorp.com

**Order Activity (Last 12 months)**:
- Total orders: 47
- Total revenue: $892,000
- Average order value: $18,978
- Order frequency: ~4 orders/month

**Product Mix**:
- 60% Electronic Components
- 25% Cables & Connectors
- 15% Mounting Hardware

**Delivery Performance** (TMS):
- On-time delivery: 89%
- Average delivery time: 4.2 days
- Exceptions: 5 (11%)

**Issues to Address**:
1. Recent exception on SHP-2024-1892 (customs hold) - still unresolved
2. 3 orders in last quarter had partial shipments due to stock

**Opportunities**:
- Their orders for SKU-SENSOR-T1 increased 40% last quarter
- May be a good candidate for volume discount on sensors

**Open Orders**:
- ORD-2024-0501: $34,500, status: delayed (customs)
- ORD-2024-0515: $12,800, status: processing
```

---

## Use Case 7: End-of-Day Operations Summary

### Scenario
Operations manager wants a daily wrap-up.

### Query

```
User: "Give me today's operations summary"

Agent: **Daily Operations Summary - January 15, 2024**

**Orders** (OMS):
- New orders received: 23
- Orders shipped: 18
- Orders pending: 47
- Revenue today: $78,450

**Shipments** (TMS):
- Dispatched: 18
- In transit: 45
- Delivered: 12
- Exceptions: 4 (see alert list)

**Inventory** (WMS):
- Stock movements: 156
- Items received: 890 units (3 POs)
- Items shipped: 678 units
- Low stock alerts: 4 SKUs

**Procurement** (SRM):
- POs created: 2
- POs received: 3
- Pending deliveries: 8

**Key Alerts**:
🔴 2 critical stockout risks (SKU-CHIP-22, SKU-CABLE-USB)
🟡 1 urgent customer order delayed (MegaCorp)
🟡 4 shipment exceptions need attention

**Tomorrow's Priorities**:
1. Resolve MegaCorp customs hold
2. Emergency PO for SKU-CABLE-USB
3. 12 orders ready for shipment
```

---

## Use Case Summary Table

| # | Use Case | Systems Queried | Value Proposition |
|---|----------|-----------------|-------------------|
| 1 | Order Delay Investigation | OMS, TMS, WMS, SRM | Root cause in seconds vs 30 min |
| 2 | Product Availability | WMS, OMS, SRM | Confident promise to customer |
| 3 | Supplier Comparison | SRM, WMS (quality) | Data-driven procurement |
| 4 | Exception Handling | TMS, OMS, WMS | Prioritized action list |
| 5 | Low Stock Analysis | WMS, SRM, OMS | Proactive vs reactive |
| 6 | Customer Summary | OMS, TMS, SRM | Meeting prep in 1 query |
| 7 | Daily Summary | All 4 systems | Executive dashboard via chat |

---

## Demo Scenarios for Prototype

For your prototype, prepare these specific scenarios with seed data:

1. **Happy path**: Order inquiry with clear status
2. **Investigation**: Delayed order with inventory + supplier chain
3. **Alert**: Morning stockout check
4. **Comparison**: Supplier evaluation
5. **Multi-turn**: Follow-up questions referencing previous answers

Each scenario should have corresponding test data in your seed script.
