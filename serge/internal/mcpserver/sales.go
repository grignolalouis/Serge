package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerSalesTools(s *mcp.StdioServer, ctrl *controller.SalesController) {
	s.RegisterTool(
		mcp.NewTool("oms_get_customer_order",
			mcp.WithDescription("Get a customer order by ID: customer, status, dates, priority, notes, and line items. Line items reference products (price comes from the product entity). Usually the first tool to call when investigating an order."),
			mcp.WithString("customer_order_id", mcp.Required(), mcp.Description("e.g. ORD-007")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetCustomerOrder(stringArg(req.Params.Arguments, "customer_order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_get_customer",
			mcp.WithDescription("Get customer profile: name, company, segment (retail/wholesale), address, delivery zone."),
			mcp.WithString("customer_id", mcp.Required(), mcp.Description("e.g. CUST-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetCustomer(stringArg(req.Params.Arguments, "customer_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_customer_orders_by_customer",
			mcp.WithDescription("List all orders placed by a specific customer."),
			mcp.WithString("customer_id", mcp.Required(), mcp.Description("e.g. CUST-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListCustomerOrdersByCustomer(stringArg(req.Params.Arguments, "customer_id")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_customer_orders_by_status",
			mcp.WithDescription("List customer orders filtered by status: pending, processing, shipped, delivered, or cancelled."),
			mcp.WithString("status", mcp.Required(), mcp.Description("pending, processing, shipped, delivered, or cancelled")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListCustomerOrdersByStatus(stringArg(req.Params.Arguments, "status")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_overdue_customer_orders",
			mcp.WithDescription("List customer orders past their required date and not yet delivered or cancelled."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListOverdueCustomerOrders())
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_all_customer_orders",
			mcp.WithDescription("List all customer orders in the system."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListAllCustomerOrders())
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_list_customers",
			mcp.WithDescription("List all customers with their profiles."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListCustomers())
		},
	)

	s.RegisterTool(
		mcp.NewTool("oms_search_customer_orders_by_product",
			mcp.WithDescription("Find all customer orders containing a specific product. Use to understand demand or trace orders affected by a stock shortage."),
			mcp.WithString("product_id", mcp.Required(), mcp.Description("e.g. PROD-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListCustomerOrdersByProduct(stringArg(req.Params.Arguments, "product_id")))
		},
	)
}
