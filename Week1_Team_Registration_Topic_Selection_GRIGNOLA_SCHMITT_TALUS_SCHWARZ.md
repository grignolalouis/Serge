# Case Study: AI Agent for Supply Chain Operations

LOUIS GRIGNOLA
GAUTHIER SCHMITT
MAEL TALUS
MATTEO SCHWARZ

**Course:** SE618-1 Logistics
**Focus:** Agentic AI applied to Supply Chain Management

## Why This Project

**Our Background:** Software Engineering, Cybersecurity, AI Engineering, and Generative AI.

**Objective:** Apply our technical expertise to the Supply Chain domain through a project that:
- Addresses a real industry problem (data silos, slow information access)
- Aligns with current trends (Agentic AI is emerging in SC, first papers 2024-2025)
- Bridges our AI skills with Supply Chain knowledge from this course
- Delivers practical value, not just theoretical analysis

This project lets us **learn Supply Chain by building for it**.

## Key Idea

Build an **AI assistant** that can query across multiple supply chain data sources and answer operational questions in natural language.

**Problem:** Supply chain data is siloed across systems (procurement, warehouse, orders, transport). Workers spend significant time manually gathering and cross-referencing information.

**Solution:** An agentic AI that understands supply chain context, queries the right sources, and synthesizes answers.

## Part 1: Supply Chain Model

### Simplified Distribution Company

```
Suppliers → Warehouse → Transport → Customers
```

### Four Data Domains

| System | What It Manages | Example Data |
|--------|-----------------|--------------|
| **SRM** | Suppliers & Purchasing | Suppliers, Purchase Orders, Lead Times |
| **WMS** | Inventory | Products, Stock Levels, Movements |
| **OMS** | Sales | Customers, Orders, Order Lines |
| **TMS** | Delivery | Shipments, Carriers, Tracking |

### Data Scale (Simulated)
- 5 suppliers
- 30 products
- 50 customers
- 500 orders (3 months history)
- 2 warehouses


## Part 2: AI Agent Architecture

### How It Works

```
┌─────────────────────────────────────────┐
│            User Question                │
│  "Why is order #123 delayed?"           │
└──────────────────┬──────────────────────┘
                   ▼
┌─────────────────────────────────────────┐
│           AI Agent (LLM)                │
│   Reasons about what data is needed     │
└──────────────────┬──────────────────────┘
                   ▼
        ┌──────────┼──────────┐
        ▼          ▼          ▼
    ┌───────┐  ┌───────┐  ┌───────┐
    │  OMS  │  │  WMS  │  │  TMS  │
    │ Tool  │  │ Tool  │  │ Tool  │
    └───────┘  └───────┘  └───────┘
        │          │          │
        └──────────┼──────────┘
                   ▼
┌─────────────────────────────────────────┐
│         Synthesized Answer              │
│  "Order delayed due to inventory        │
│   shortage. PO arriving Jan 18."        │
└─────────────────────────────────────────┘
```

### Agent Capabilities

1. **Query** - Access all 4 data systems
2. **Reason** - Determine which sources to check based on the question can do long range work so reasoning and wheck lmany source, launch exploratory worker subagent etc.
3. **Synthesize** - Combine data into actionable answers
4. **Recommend** - Suggest next steps based on analysis

## Technology Stack

| Component | Technology |
|-----------|------------|
| Backend | Go (Fiber) |
| Agent Framework | tRPC-Agent-Go |
| Database | PostgreSQL |
| Frontend | Next.js |
