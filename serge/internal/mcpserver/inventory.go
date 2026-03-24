package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerInventoryTools(s *mcp.StdioServer, ctrl *controller.InventoryController) {

	s.RegisterTool(
		mcp.NewTool("wms_get_product",
			mcp.WithDescription("Get product details by ID: name, SKU, category, price, weight, reorder point. Use when you have a product ID and need its details."),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("Product ID, e.g. PROD-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetProduct(stringArg(req.Params.Arguments, "product_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_get_product_by_sku",
			mcp.WithDescription("Look up a product by its SKU code. Use when you have a SKU (e.g. MANGO-NDM-5) instead of a product ID."),
			mcp.WithString("sku", mcp.Required(), mcp.Description("Product SKU, e.g. MANGO-NDM-5")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetProductBySKU(stringArg(req.Params.Arguments, "sku"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_list_products",
			mcp.WithDescription("List all products in the catalog with their details. Use to discover available products."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListProducts())
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_check_inventory",
			mcp.WithDescription("Check current inventory level for a product: quantity on hand, reserved, available, and whether it needs reordering. Use when investigating stock availability for an order."),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("Product ID, e.g. PROD-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.CheckInventory(stringArg(req.Params.Arguments, "product_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_list_low_stock",
			mcp.WithDescription("List all products currently below their reorder point. Use to identify stock shortages that may be causing order delays."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListLowStock())
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_get_stock_movements",
			mcp.WithDescription("Get the history of stock movements (inbound/outbound) for a product. Each movement references the source PO or sales order. Use to understand how stock was consumed or replenished."),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("Product ID, e.g. PROD-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.GetStockMovements(stringArg(req.Params.Arguments, "product_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_list_all_inventory",
			mcp.WithDescription("Get a full inventory snapshot: all products with current stock levels, availability, and reorder status. Use for a complete warehouse overview."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListAllInventory())
		},
	)

	s.RegisterTool(
		mcp.NewTool("wms_get_warehouse",
			mcp.WithDescription("Get warehouse details: name and location. Use when you need warehouse context."),
			mcp.WithString("warehouse_id", mcp.Required(), mcp.Description("Warehouse ID, e.g. WH-BKK")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetWarehouse(stringArg(req.Params.Arguments, "warehouse_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)
}
