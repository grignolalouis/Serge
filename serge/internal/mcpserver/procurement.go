package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerProcurementTools(s *mcp.StdioServer, ctrl *controller.ProcurementController) {

	s.RegisterTool(
		mcp.NewTool("srm_get_supplier",
			mcp.WithDescription("Get a supplier's profile: name, region, contact, lead time, and rating. Use when you need details about a specific supplier."),
			mcp.WithString("supplier_id", mcp.Required(), mcp.Description("Supplier ID, e.g. SUP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetSupplier(stringArg(req.Params.Arguments, "supplier_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_suppliers",
			mcp.WithDescription("List all suppliers with their profiles. Use to discover available suppliers or before comparing them."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListSuppliers())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_get_purchase_order",
			mcp.WithDescription("Get a purchase order's details: supplier, status, dates, line items. Use when investigating a specific PO."),
			mcp.WithString("purchase_order_id", mcp.Required(), mcp.Description("Purchase order ID, e.g. PO-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetPurchaseOrder(stringArg(req.Params.Arguments, "purchase_order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_purchase_orders_by_supplier",
			mcp.WithDescription("List all purchase orders for a given supplier. Use to check what has been ordered from a supplier and delivery history."),
			mcp.WithString("supplier_id", mcp.Required(), mcp.Description("Supplier ID, e.g. SUP-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListPurchaseOrdersBySupplier(stringArg(req.Params.Arguments, "supplier_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_purchase_orders_by_status",
			mcp.WithDescription("List purchase orders filtered by status. Use to find pending, shipped, or overdue POs."),
			mcp.WithString("status", mcp.Required(), mcp.Description("PO status: pending, confirmed, shipped, received, or cancelled")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListPurchaseOrdersByStatus(stringArg(req.Params.Arguments, "status")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_overdue_purchase_orders",
			mcp.WithDescription("List purchase orders that are past their expected delivery date and not yet received. Use to identify supplier delays."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOverduePurchaseOrders())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_compare_suppliers",
			mcp.WithDescription("Compare all suppliers side by side: total POs, on-time delivery rate, average lead time. Use when evaluating supplier performance or choosing a supplier."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.CompareSuppliers())
		},
	)

	s.RegisterTool(
		mcp.NewTool("srm_list_all_purchase_orders",
			mcp.WithDescription("List all purchase orders across all suppliers with their status, dates, and line items. Use for full procurement visibility."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListAllPurchaseOrders())
		},
	)
}
