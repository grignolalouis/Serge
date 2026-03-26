package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerProcurementTools(s *mcp.StdioServer, ctrl *controller.ProcurementController) {
	s.RegisterTool(
		mcp.NewTool("srm_get_supplier",
			mcp.WithDescription("Get a supplier's profile: name, region, contact, lead time, rating, product categories, payment terms."),
			mcp.WithString("supplier_id", mcp.Required(), mcp.Description("e.g. SUP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetSupplier(stringArg(req.Params.Arguments, "supplier_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_suppliers",
			mcp.WithDescription("List all suppliers."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSuppliers())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_get_supplier_product",
			mcp.WithDescription("Get a supplier product by ID. A supplier product is an item a supplier sells, with the supplier's unit cost. Use to find the purchase price for a product from a specific supplier."),
			mcp.WithString("supplier_product_id", mcp.Required(), mcp.Description("e.g. SP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetSupplierProduct(stringArg(req.Params.Arguments, "supplier_product_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_supplier_products_by_supplier",
			mcp.WithDescription("List all products a specific supplier offers, with their unit costs."),
			mcp.WithString("supplier_id", mcp.Required(), mcp.Description("e.g. SUP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSupplierProductsBySupplier(stringArg(req.Params.Arguments, "supplier_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_supplier_products_by_product",
			mcp.WithDescription("Find all suppliers that sell a given product, with their unit costs. Use to compare sourcing options."),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("Our product ID, e.g. PROD-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSupplierProductsByProduct(stringArg(req.Params.Arguments, "product_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_get_supplier_order",
			mcp.WithDescription("Get a supplier order (purchase order) by ID: supplier, status, dates, line items. Line items reference supplier products."),
			mcp.WithString("supplier_order_id", mcp.Required(), mcp.Description("e.g. PO-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetSupplierOrder(stringArg(req.Params.Arguments, "supplier_order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_all_supplier_orders",
			mcp.WithDescription("List all supplier orders (purchase orders) with status, dates, and line items."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListAllSupplierOrders())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_supplier_orders_by_supplier",
			mcp.WithDescription("List all supplier orders placed with a specific supplier."),
			mcp.WithString("supplier_id", mcp.Required(), mcp.Description("e.g. SUP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSupplierOrdersBySupplier(stringArg(req.Params.Arguments, "supplier_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_supplier_orders_by_status",
			mcp.WithDescription("List supplier orders filtered by status: pending, confirmed, shipped, received, or cancelled."),
			mcp.WithString("status", mcp.Required(), mcp.Description("pending, confirmed, shipped, received, or cancelled")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSupplierOrdersByStatus(stringArg(req.Params.Arguments, "status")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_overdue_supplier_orders",
			mcp.WithDescription("List supplier orders past expected delivery and not yet received."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOverdueSupplierOrders())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_compare_suppliers",
			mcp.WithDescription("Compare all suppliers: total orders, on-time rate, average lead time. Use for supplier evaluation."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.CompareSuppliers())
		},
	)
}
