package main

import (
	"log"
	"os"

	"github.com/serge-music/serge/internal/controller"
	"github.com/serge-music/serge/internal/mcpserver"
	"github.com/serge-music/serge/internal/repository"
	"github.com/serge-music/serge/internal/seed"
	"github.com/serge-music/serge/internal/service"
)

func main() {
	procRepo := repository.NewProcurementRepository()
	invRepo := repository.NewInventoryRepository()
	salesRepo := repository.NewSalesRepository()
	shipRepo := repository.NewShippingRepository()

	dataDir := "data"
	if len(os.Args) > 1 {
		dataDir = os.Args[1]
	}
	if err := seed.LoadAll(dataDir, procRepo, invRepo, salesRepo, shipRepo); err != nil {
		log.Fatalf("seed: %v", err)
	}

	server := mcpserver.New(
		controller.NewProcurementController(service.NewProcurementService(procRepo)),
		controller.NewInventoryController(service.NewInventoryService(invRepo)),
		controller.NewSalesController(service.NewSalesService(salesRepo)),
		controller.NewShippingController(service.NewShippingService(shipRepo)),
	)

	log.Println("SERGE MCP server starting on stdio...")
	server.Start()
}
