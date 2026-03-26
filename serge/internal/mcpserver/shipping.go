package mcpserver

import (
	"context"

	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

func registerShippingTools(s *mcp.StdioServer, ctrl *controller.ShippingController) {
	s.RegisterTool(
		mcp.NewTool("tms_get_shipment",
			mcp.WithDescription("Get shipment details including carrier info and tracking events."),
			mcp.WithString("shipment_id", mcp.Required(), mcp.Description("e.g. SHP-005")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetShipment(stringArg(req.Params.Arguments, "shipment_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_track_customer_order",
			mcp.WithDescription("Find all shipments for a customer order and return tracking details. Returns an error if no shipment exists yet. An order may have multiple shipments (split shipment)."),
			mcp.WithString("customer_order_id", mcp.Required(), mcp.Description("e.g. ORD-005")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.TrackCustomerOrder(stringArg(req.Params.Arguments, "customer_order_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_list_exception_shipments",
			mcp.WithDescription("List all shipments with delivery exceptions."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListExceptionShipments())
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_get_carrier",
			mcp.WithDescription("Get carrier profile: name, type (ground/refrigerated), cost per kg, average transit time."),
			mcp.WithString("carrier_id", mcp.Required(), mcp.Description("e.g. CAR-001")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			r, err := ctrl.GetCarrier(stringArg(req.Params.Arguments, "carrier_id"))
			if err != nil {
				return errResult(err)
			}
			return jsonResult(r)
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_list_shipments_by_status",
			mcp.WithDescription("List shipments by status: pending, in_transit, delivered, or exception."),
			mcp.WithString("status", mcp.Required(), mcp.Description("pending, in_transit, delivered, or exception")),
		),
		func(_ context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListShipmentsByStatus(stringArg(req.Params.Arguments, "status")))
		},
	)

	s.RegisterTool(
		mcp.NewTool("tms_list_carriers",
			mcp.WithDescription("List all carriers with their profiles."),
		),
		func(_ context.Context, _ *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return jsonResult(ctrl.ListCarriers())
		},
	)
}
