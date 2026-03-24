package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerShippingTools(s *mcp.StdioServer, ctrl *controller.ShippingController) {

	s.RegisterTool(
		mcp.NewTool("tms_get_shipment",
			mcp.WithDescription("Get shipment details including carrier info and full tracking event history. Use when you have a shipment ID."),
			mcp.WithString("shipment_id", mcp.Required(), mcp.Description("Shipment ID, e.g. SHP-005")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetShipment(stringArg(req.Params.Arguments, "shipment_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_track_order",
			mcp.WithDescription("Find the shipment associated with an order and return its full tracking details (carrier, status, events). Use when you have an order ID and want to know its delivery status. Returns an error if no shipment exists yet for the order."),
			mcp.WithString("order_id", mcp.Required(), mcp.Description("Order ID, e.g. ORD-005")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.TrackOrder(stringArg(req.Params.Arguments, "order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_list_exception_shipments",
			mcp.WithDescription("List all shipments with delivery exceptions (e.g. vehicle breakdown, address issue). Use to identify shipments that need intervention."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListExceptionShipments())
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_get_carrier",
			mcp.WithDescription("Get carrier profile: name, type (ground/refrigerated), cost per kg, and average transit time. Use to understand carrier capabilities."),
			mcp.WithString("carrier_id", mcp.Required(), mcp.Description("Carrier ID, e.g. CAR-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := ctrl.GetCarrier(stringArg(req.Params.Arguments, "carrier_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(result)
		},
	)
}
