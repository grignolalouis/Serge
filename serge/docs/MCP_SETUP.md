# SERGE MCP Server - Setup Guide

## Prerequisites

- **Go 1.23+** installed
- **Claude Code CLI** installed
- This repository cloned locally

## Project Structure

```
serge/
├── cmd/server/main.go          # Entry point
├── internal/
│   ├── mcpserver/              # MCP tool definitions (stdio-based)
│   ├── controller/             # Business logic controllers
│   ├── service/                # Service layer
│   ├── repository/             # In-memory data repositories
│   └── seed/                   # JSON seed data loader
├── data/                       # 11 JSON seed files
├── bin/                        # Compiled binary output
├── .mcp.json                   # MCP server config for Claude Code
└── .claude/settings.local.json # Enables the MCP server
```

## 1. Build the MCP Server

```bash
cd <PROJECT_ROOT>/serge
go build -o bin/serge-mcp ./cmd/server/
```

## 2. Configure `.mcp.json`

Create or update `.mcp.json` at the repo root:

```json
{
  "mcpServers": {
    "serge-supply-chain": {
      "command": "<PROJECT_ROOT>/serge/bin/serge-mcp",
      "args": ["<PROJECT_ROOT>/serge/data"]
    }
  }
}
```

Replace `<PROJECT_ROOT>` with the absolute path to your `SE618-1_Logistics_Case_Study_Analysis` directory.

Example:
```
/home/user/projects/SE618-1_Logistics_Case_Study_Analysis
```

## 3. Configure `.claude/settings.local.json`

Create `.claude/settings.local.json` in the repo root:

```json
{
  "enabledMcpjsonServers": [
    "serge-supply-chain"
  ]
}
```

## 4. Verify

```bash
# Test the binary runs and loads seed data
<PROJECT_ROOT>/serge/bin/serge-mcp <PROJECT_ROOT>/serge/data
# Should print: "SERGE MCP server starting on stdio..."
# Ctrl+C to stop

# Then launch Claude Code from the serge/ directory
cd <PROJECT_ROOT>/serge
claude
```

Claude Code will automatically detect `.mcp.json` and connect to the MCP server. You should see 21 tools available under the `serge-supply-chain` prefix.

## Available MCP Tools (21 total)

### SRM - Supplier Relationship Management (7 tools)
| Tool | Description |
|------|-------------|
| `srm_get_supplier` | Get supplier details by ID |
| `srm_list_suppliers` | List all suppliers |
| `srm_get_purchase_order` | Get purchase order details |
| `srm_list_purchase_orders_by_supplier` | POs filtered by supplier |
| `srm_list_purchase_orders_by_status` | POs filtered by status |
| `srm_list_overdue_purchase_orders` | List overdue POs |
| `srm_compare_suppliers` | Compare all suppliers side-by-side |

### WMS - Warehouse Management (6 tools)
| Tool | Description |
|------|-------------|
| `wms_get_product` | Get product by ID |
| `wms_get_product_by_sku` | Get product by SKU |
| `wms_list_products` | List all products |
| `wms_check_inventory` | Check stock levels for a product |
| `wms_list_low_stock` | List products below reorder point |
| `wms_get_stock_movements` | Get inbound/outbound movements |

### OMS - Order Management (5 tools)
| Tool | Description |
|------|-------------|
| `oms_get_order` | Get full order details |
| `oms_get_customer` | Get customer details |
| `oms_list_orders_by_customer` | Orders filtered by customer |
| `oms_list_orders_by_status` | Orders filtered by status |
| `oms_list_overdue_orders` | List overdue orders |

### TMS - Transport Management (4 tools)
| Tool | Description |
|------|-------------|
| `tms_get_shipment` | Get shipment details |
| `tms_track_order` | Track shipment for an order |
| `tms_list_exception_shipments` | List shipments with exceptions |
| `tms_get_carrier` | Get carrier details |

## Seed Data Summary

| File | Records | Description |
|------|---------|-------------|
| `suppliers.json` | 3 | Thai mango farms |
| `purchase_orders.json` | ~10 | Procurement orders to suppliers |
| `products.json` | 5 | Mango product SKUs |
| `warehouses.json` | 1 | Bangkok central warehouse |
| `inventory.json` | 5 | Stock levels per product |
| `stock_movements.json` | ~20 | Inbound/outbound records |
| `customers.json` | 50 | Thai B2B customers |
| `orders.json` | 500 | Customer sales orders |
| `carriers.json` | 2 | Kerry Express + SCG Cold Chain |
| `shipments.json` | 400 | Delivery shipments |
| `tracking_events.json` | ~1000 | Shipment tracking history |
