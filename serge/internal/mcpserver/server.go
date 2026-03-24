package mcpserver

import (
	mcp "trpc.group/trpc-go/trpc-mcp-go"

	"github.com/serge-music/serge/internal/controller"
)

// Server wraps the MCP stdio server and registers all supply chain tools.
type Server struct {
	stdio *mcp.StdioServer
}

// New creates an MCP server with all tools registered from the four controllers.
func New(
	procurement *controller.ProcurementController,
	inventory *controller.InventoryController,
	sales *controller.SalesController,
	shipping *controller.ShippingController,
) *Server {
	s := &Server{
		stdio: mcp.NewStdioServer("serge-supply-chain", "1.0.0"),
	}

	registerProcurementTools(s.stdio, procurement)
	registerInventoryTools(s.stdio, inventory)
	registerSalesTools(s.stdio, sales)
	registerShippingTools(s.stdio, shipping)

	return s
}

// Start blocks and serves the MCP protocol over stdio.
func (s *Server) Start() {
	s.stdio.Start()
}
