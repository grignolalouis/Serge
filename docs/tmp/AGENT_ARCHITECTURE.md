# Agent Architecture Design

## Overview

A single orchestrator agent with **specialized tools** for each data source. This is simpler than multi-agent while still demonstrating cross-source reasoning.

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              NEXT.JS FRONTEND                           │
│                         (Chat UI + Dashboard)                           │
└─────────────────────────────────┬───────────────────────────────────────┘
                                  │ SSE / WebSocket
                                  ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                           GO BACKEND (Fiber)                            │
│  ┌──────────────────────────────────────────────────────────────────┐  │
│  │                     tRPC-Agent-Go Runner                          │  │
│  │  ┌────────────────────────────────────────────────────────────┐  │  │
│  │  │                   SUPPLY CHAIN AGENT                        │  │  │
│  │  │                                                             │  │  │
│  │  │   ┌─────────────────────────────────────────────────────┐  │  │  │
│  │  │   │                    LLM (Claude)                      │  │  │  │
│  │  │   │         System Prompt + Tool Definitions            │  │  │  │
│  │  │   └─────────────────────────────────────────────────────┘  │  │  │
│  │  │                          │                                  │  │  │
│  │  │          ┌───────────────┼───────────────┐                 │  │  │
│  │  │          ▼               ▼               ▼                 │  │  │
│  │  │   ┌──────────┐    ┌──────────┐    ┌──────────┐            │  │  │
│  │  │   │ SRM Tool │    │ WMS Tool │    │ OMS Tool │            │  │  │
│  │  │   └──────────┘    └──────────┘    └──────────┘            │  │  │
│  │  │          │               │               │                 │  │  │
│  │  │          ▼               ▼               ▼                 │  │  │
│  │  │   ┌──────────┐    ┌──────────┐    ┌──────────┐            │  │  │
│  │  │   │ TMS Tool │    │ Analysis │    │ Memory   │            │  │  │
│  │  │   └──────────┘    │   Tool   │    │  Tool    │            │  │  │
│  │  │                   └──────────┘    └──────────┘            │  │  │
│  │  └────────────────────────────────────────────────────────────┘  │  │
│  └──────────────────────────────────────────────────────────────────┘  │
│                                  │                                      │
└──────────────────────────────────┼──────────────────────────────────────┘
                                   │
                    ┌──────────────┼──────────────┐
                    ▼              ▼              ▼
              ┌──────────┐  ┌──────────┐  ┌──────────┐
              │   SRM    │  │   WMS    │  │   OMS    │
              │   DB     │  │   DB     │  │   DB     │
              └──────────┘  └──────────┘  └──────────┘
                                   │
                              ┌────┴────┐
                              │   TMS   │
                              │   DB    │
                              └─────────┘

        (Can be single PostgreSQL with separate schemas)
```

---

## 1. Agent Design

### Single Agent with Tool Selection

The agent receives a user query and decides which tools to call (possibly multiple in sequence).

```go
// Simplified conceptual structure
type SupplyChainAgent struct {
    llm       LLMClient
    tools     []Tool
    memory    MemoryStore
    systemPrompt string
}

func (a *SupplyChainAgent) Process(ctx context.Context, userQuery string) (string, error) {
    // 1. Agent reasons about what info is needed
    // 2. Agent calls appropriate tools
    // 3. Agent synthesizes response from tool outputs
    // 4. Agent may iterate if more info needed
}
```

### System Prompt Design

```markdown
You are a Supply Chain Assistant helping warehouse managers, planners, and logistics coordinators.

You have access to 4 data systems:
- SRM: Supplier information, purchase orders, supplier performance
- WMS: Product catalog, inventory levels, stock movements
- OMS: Customers, sales orders, order status
- TMS: Shipments, carriers, delivery tracking

When answering questions:
1. Identify which systems contain relevant data
2. Query the appropriate tools
3. Cross-reference information across systems when needed
4. Provide clear, actionable answers with specific data

Always cite which system the data came from.
If you cannot find the requested information, say so clearly.
```

---

## 2. Tool Definitions

### Tool 1: SRM (Supplier Relationship Management)

```go
// Tool: query_srm
// Description: Query supplier and purchase order data

type SRMQueryInput struct {
    QueryType string `json:"query_type"` // "supplier_info", "purchase_orders", "supplier_performance"

    // For supplier_info
    SupplierID   string `json:"supplier_id,omitempty"`
    SupplierName string `json:"supplier_name,omitempty"`

    // For purchase_orders
    POID        string `json:"po_id,omitempty"`
    ProductID   string `json:"product_id,omitempty"`
    Status      string `json:"status,omitempty"`      // pending, confirmed, shipped, received
    DateFrom    string `json:"date_from,omitempty"`
    DateTo      string `json:"date_to,omitempty"`

    // For supplier_performance
    SupplierIDs []string `json:"supplier_ids,omitempty"` // Compare multiple
}

type SRMQueryOutput struct {
    QueryType string      `json:"query_type"`
    Results   interface{} `json:"results"`
    Count     int         `json:"count"`
}
```

**Example Queries:**
- "Get supplier info for supplier_id=ABC"
- "List pending purchase orders for product SKU-123"
- "Compare performance of suppliers X, Y, Z"

---

### Tool 2: WMS (Warehouse Management)

```go
// Tool: query_wms
// Description: Query product and inventory data

type WMSQueryInput struct {
    QueryType string `json:"query_type"` // "product_info", "inventory", "stock_movements", "low_stock"

    // For product_info
    ProductID string `json:"product_id,omitempty"`
    SKU       string `json:"sku,omitempty"`
    Category  string `json:"category,omitempty"`

    // For inventory
    WarehouseID string `json:"warehouse_id,omitempty"` // WH-A, WH-B, or empty for all

    // For stock_movements
    MovementType string `json:"movement_type,omitempty"` // inbound, outbound, adjustment
    DateFrom     string `json:"date_from,omitempty"`
    DateTo       string `json:"date_to,omitempty"`

    // For low_stock
    BelowReorderPoint bool `json:"below_reorder_point,omitempty"`
}

type WMSQueryOutput struct {
    QueryType string      `json:"query_type"`
    Results   interface{} `json:"results"`
    Count     int         `json:"count"`
}
```

**Example Queries:**
- "Get inventory for product SKU-456 across all warehouses"
- "List products below reorder point"
- "Show stock movements for last 7 days"

---

### Tool 3: OMS (Order Management)

```go
// Tool: query_oms
// Description: Query customer and order data

type OMSQueryInput struct {
    QueryType string `json:"query_type"` // "customer_info", "order_details", "order_search", "order_lines"

    // For customer_info
    CustomerID   string `json:"customer_id,omitempty"`
    CustomerName string `json:"customer_name,omitempty"`

    // For order_details
    OrderID string `json:"order_id,omitempty"`

    // For order_search
    Status     string `json:"status,omitempty"`     // pending, processing, shipped, delivered
    Priority   string `json:"priority,omitempty"`   // urgent, high, normal, low
    CustomerID string `json:"customer_id,omitempty"`
    DateFrom   string `json:"date_from,omitempty"`
    DateTo     string `json:"date_to,omitempty"`

    // For order_lines
    ProductID string `json:"product_id,omitempty"` // Orders containing this product
}

type OMSQueryOutput struct {
    QueryType string      `json:"query_type"`
    Results   interface{} `json:"results"`
    Count     int         `json:"count"`
}
```

**Example Queries:**
- "Get details for order #12345"
- "List all urgent orders not yet shipped"
- "Which orders contain product SKU-789?"

---

### Tool 4: TMS (Transport Management)

```go
// Tool: query_tms
// Description: Query shipment and carrier data

type TMSQueryInput struct {
    QueryType string `json:"query_type"` // "shipment_details", "shipment_search", "carrier_info", "exceptions"

    // For shipment_details
    ShipmentID     string `json:"shipment_id,omitempty"`
    TrackingNumber string `json:"tracking_number,omitempty"`
    OrderID        string `json:"order_id,omitempty"`

    // For shipment_search
    Status     string `json:"status,omitempty"` // in_transit, delivered, exception
    CarrierID  string `json:"carrier_id,omitempty"`
    DateFrom   string `json:"date_from,omitempty"`
    DateTo     string `json:"date_to,omitempty"`

    // For carrier_info
    CarrierID string `json:"carrier_id,omitempty"`

    // For exceptions
    // Returns all shipments with exception status
}

type TMSQueryOutput struct {
    QueryType string      `json:"query_type"`
    Results   interface{} `json:"results"`
    Count     int         `json:"count"`
}
```

**Example Queries:**
- "Track shipment for order #12345"
- "List all shipments with exceptions"
- "Show carrier performance metrics"

---

### Tool 5: Analysis (Cross-Source)

```go
// Tool: analyze_supply_chain
// Description: Perform cross-source analysis and calculations

type AnalysisInput struct {
    AnalysisType string `json:"analysis_type"`
    // Types: "order_delay_root_cause", "product_availability", "supplier_comparison", "demand_trend"

    // Parameters vary by analysis type
    OrderID     string   `json:"order_id,omitempty"`
    ProductID   string   `json:"product_id,omitempty"`
    SupplierIDs []string `json:"supplier_ids,omitempty"`
    DateRange   int      `json:"date_range_days,omitempty"`
}
```

This tool internally calls multiple other tools and synthesizes results.

---

## 3. Agent Reasoning Flow

### Example: "Why is order #ORD-123 delayed?"

```
User: "Why is order #ORD-123 delayed?"

Agent thinking:
1. Need order details → call query_oms(order_details, ORD-123)
2. Check shipment status → call query_tms(shipment_details, order_id=ORD-123)
3. If not shipped, check inventory → call query_wms(inventory, product_ids from order)
4. If inventory issue, check purchase orders → call query_srm(purchase_orders, product_id)

Tool calls:
├── query_oms: Order #ORD-123, status=processing, required_date=2024-01-15
│   └── Order lines: SKU-456 (qty: 50), SKU-789 (qty: 20)
│
├── query_tms: No shipment found for ORD-123
│
├── query_wms: SKU-456 available=10, SKU-789 available=25
│   └── SKU-456 is short by 40 units
│
└── query_srm: PO-789 for SKU-456, status=shipped, expected=2024-01-18
    └── Supplier: Acme Corp, current lead time running 3 days late

Agent response:
"Order #ORD-123 is delayed because:

1. **Inventory shortage**: SKU-456 needs 50 units but only 10 are available
2. **Pending replenishment**: PO-789 from Acme Corp is in transit, expected Jan 18
3. **Supplier delay**: Acme Corp is running 3 days behind their standard lead time

**Recommendation**: The order can ship once PO-789 arrives. Expected ship date: Jan 19.
Consider expedited shipping to meet the Jan 15 required date."
```

---

## 4. Memory Design

### Conversation Memory
- Store recent conversation history per session
- Allow agent to reference previous queries

### Entity Memory (Optional Enhancement)
- Track entities mentioned in conversation
- "The customer I mentioned earlier" → resolve to specific customer ID

```go
type ConversationMemory struct {
    SessionID    string
    Messages     []Message
    MentionedEntities map[string]string // "the order" -> "ORD-123"
    CreatedAt    time.Time
    LastAccess   time.Time
}
```

---

## 5. API Design (Go/Fiber)

### Endpoints

```go
// POST /api/chat - Main chat endpoint
type ChatRequest struct {
    SessionID string `json:"session_id"`
    Message   string `json:"message"`
}

type ChatResponse struct {
    SessionID string `json:"session_id"`
    Response  string `json:"response"`
    ToolsUsed []string `json:"tools_used"` // For transparency
    Sources   []DataSource `json:"sources"` // Which systems were queried
}

// GET /api/chat/stream/:session_id - SSE endpoint for streaming
// Streams: thinking, tool_calls, response chunks

// GET /api/sessions/:session_id/history - Get conversation history

// POST /api/sessions - Create new session

// DELETE /api/sessions/:session_id - End session
```

### SSE Events

```go
type SSEEvent struct {
    Type    string      `json:"type"` // "thinking", "tool_call", "tool_result", "response", "done"
    Content interface{} `json:"content"`
}

// Example stream:
// {"type": "thinking", "content": "Checking order details..."}
// {"type": "tool_call", "content": {"tool": "query_oms", "input": {...}}}
// {"type": "tool_result", "content": {"tool": "query_oms", "output": {...}}}
// {"type": "response", "content": "Order #ORD-123 is delayed because..."}
// {"type": "done", "content": null}
```

---

## 6. Project Structure

```
supply-chain-agent/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── agent/
│   │   ├── agent.go             # Agent setup with tRPC-Agent-Go
│   │   ├── prompts.go           # System prompts
│   │   └── memory.go            # Memory management
│   ├── tools/
│   │   ├── srm.go               # SRM tool implementation
│   │   ├── wms.go               # WMS tool implementation
│   │   ├── oms.go               # OMS tool implementation
│   │   ├── tms.go               # TMS tool implementation
│   │   └── analysis.go          # Cross-source analysis tool
│   ├── db/
│   │   ├── postgres.go          # DB connection
│   │   ├── queries/             # SQL queries per system
│   │   │   ├── srm.sql
│   │   │   ├── wms.sql
│   │   │   ├── oms.sql
│   │   │   └── tms.sql
│   │   └── models/              # Data models
│   │       ├── supplier.go
│   │       ├── product.go
│   │       ├── order.go
│   │       └── shipment.go
│   ├── api/
│   │   ├── routes.go            # Fiber routes
│   │   ├── handlers.go          # Request handlers
│   │   └── middleware.go        # Auth, logging, etc.
│   └── config/
│       └── config.go            # Configuration
├── scripts/
│   ├── seed_data.go             # Generate fake data
│   └── migrate.sql              # DB schema
├── frontend/                    # Next.js app (separate or monorepo)
│   ├── app/
│   ├── components/
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

---

## 7. Technology Choices Summary

| Component | Choice | Reason |
|-----------|--------|--------|
| Backend | Go + Fiber | Your preference, fast, good for SSE |
| Agent Framework | tRPC-Agent-Go | Tool orchestration, memory, Go native |
| LLM | Claude API | Strong reasoning, good tool use |
| Database | PostgreSQL | Single DB, multiple schemas for "systems" |
| Frontend | Next.js | Your preference |
| Streaming | SSE | Simpler than WebSocket for this use case |
