package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerSalesTools(s *mcp.StdioServer, ctrl *controller.SalesController) {

	s.RegisterTool(
		mcp.NewTool("oms_get_order",
			mcp.WithDescription("Get full order details: customer, status, dates, priority, and all line items with quantities and prices. This is usually the first tool to call when investigating an order."),
			mcp.WithString("order_id", mcp.Required(), mcp.Description("Order ID, e.g. ORD-007")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetOrder(stringArg(req.Params.Arguments, "order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_get_customer",
			mcp.WithDescription("Get customer profile: name, company, segment, and region. Use when you need customer context for an order."),
			mcp.WithString("customer_id", mcp.Required(), mcp.Description("Customer ID, e.g. CUST-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetCustomer(stringArg(req.Params.Arguments, "customer_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_orders_by_customer",
			mcp.WithDescription("List all orders placed by a specific customer. Use to see a customer's order history."),
			mcp.WithString("customer_id", mcp.Required(), mcp.Description("Customer ID, e.g. CUST-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOrdersByCustomer(stringArg(req.Params.Arguments, "customer_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_orders_by_status",
			mcp.WithDescription("List orders filtered by status. Use to find all pending, processing, shipped, or delivered orders."),
			mcp.WithString("status", mcp.Required(), mcp.Description("Order status: pending, processing, shipped, delivered, or cancelled")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOrdersByStatus(stringArg(req.Params.Arguments, "status")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_overdue_orders",
			mcp.WithDescription("List orders that are past their required delivery date and not yet delivered or cancelled. Use to identify orders that need immediate attention."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOverdueOrders())
		},
	)
}
