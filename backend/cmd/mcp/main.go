package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"serge/backend/internal"
	"serge/backend/internal/controllers"
	"serge/backend/internal/repositories"
	"serge/backend/internal/services"
	"serge/backend/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (ignore if not found)
	_ = godotenv.Load()

	// Load configuration
	cfg, err := internal.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := database.Connect(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	omsRepo := repositories.NewOmsRepository(db)
	tmsRepo := repositories.NewTmsRepository(db)
	wmsRepo := repositories.NewWmsRepository(db)
	srmRepo := repositories.NewSrmRepository(db)

	// Initialize services
	omsService := services.NewOmsService(omsRepo)
	tmsService := services.NewTmsService(tmsRepo)
	wmsService := services.NewWmsService(wmsRepo)
	srmService := services.NewSrmService(srmRepo)

	// Initialize controllers
	omsRouter := &controllers.OmsRouter{Service: omsService}
	tmsRouter := &controllers.TmsRouter{Service: tmsService}
	wmsRouter := &controllers.WmsRouter{Service: wmsService}
	srmRouter := &controllers.SrmRouter{Service: srmService}

	// Log to stderr (Claude Code captures stdout for JSON-RPC)
	fmt.Fprintf(os.Stderr, "✅ SERGE MCP Server initialized (stdio mode)\n")
	fmt.Fprintf(os.Stderr, "📦 Available tools: oms_get_order, tms_get_shipment_by_order, wms_get_inventory_by_sku, srm_get_suppliers_by_category, srm_get_purchase_orders_by_sku\n")

	// Create MCP handler
	handler := NewMCPHandler(omsRouter, tmsRouter, wmsRouter, srmRouter)

	// Run stdio server - read JSON-RPC requests from stdin, write responses to stdout
	scanner := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		var req MCPRequest
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
			response := MCPResponse{
				JSONRPC: "2.0",
				Error: &MCPError{
					Code:    -32700,
					Message: "Parse error",
				},
				ID: nil,
			}
			encoder.Encode(response)
			writer.Flush()
			continue
		}

		// Handle the request
		response := handler.HandleRequest(&req)
		encoder.Encode(response)
		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
	}
}