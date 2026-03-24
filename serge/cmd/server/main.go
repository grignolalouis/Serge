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
	// Repositories
	procRepo := repository.NewProcurementRepository()
	invRepo := repository.NewInventoryRepository()
	salesRepo := repository.NewSalesRepository()
	shipRepo := repository.NewShippingRepository()

	// Load seed data
	dataDir := "data"
	if len(os.Args) > 1 {
		dataDir = os.Args[1]
	}
	if err := seed.LoadAll(dataDir, procRepo, invRepo, salesRepo, shipRepo); err != nil {
		log.Fatalf("failed to load seed data: %v", err)
	}

	// Services
	procSvc := service.NewProcurementService(procRepo)
	invSvc := service.NewInventoryService(invRepo)
	salesSvc := service.NewSalesService(salesRepo)
	shipSvc := service.NewShippingService(shipRepo)

	// Controllers
	procCtrl := controller.NewProcurementController(procSvc)
	invCtrl := controller.NewInventoryController(invSvc)
	salesCtrl := controller.NewSalesController(salesSvc)
	shipCtrl := controller.NewShippingController(shipSvc)

	// MCP server — registers all tools and serves over stdio
	server := mcpserver.New(procCtrl, invCtrl, salesCtrl, shipCtrl)

	log.Println("SERGE MCP server starting on stdio...")
	server.Start()
}
