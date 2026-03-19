package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serge/backend/internal"
	"serge/backend/internal/controllers"
	"serge/backend/internal/repositories"
	"serge/backend/internal/services"
	"serge/backend/pkg/database"
	"syscall"

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

	// Log successful initialization
	fmt.Println("✅ SERGE Backend initialized successfully")
	fmt.Printf("🚀 Server running on %s:%s\n", cfg.Server.Host, cfg.Server.Port)
	fmt.Println("📦 Available routers: OMS, TMS, WMS, SRM")

	// Verify all routers are initialized
	_ = omsRouter
	_ = tmsRouter
	_ = wmsRouter
	_ = srmRouter

	// Setup HTTP routes
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/oms/order", handleOmsGetOrder(omsRouter))
	mux.HandleFunc("/tms/shipment", handleTmsGetShipment(tmsRouter))
	mux.HandleFunc("/wms/inventory", handleWmsGetInventory(wmsRouter))
	mux.HandleFunc("/srm/suppliers", handleSrmGetSuppliers(srmRouter))
	mux.HandleFunc("/srm/purchase-orders", handleSrmGetPurchaseOrders(srmRouter))

	// Start server in a goroutine
	go func() {
		addr := cfg.Server.Host + ":" + cfg.Server.Port
		fmt.Printf("📡 HTTP server listening on http://%s\n", addr)
		if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n🛑 Shutting down server...")
}

// Health check endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy"}`)
}

// OMS endpoints
func handleOmsGetOrder(router *controllers.OmsRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		orderNumber := r.URL.Query().Get("orderNumber")
		if orderNumber == "" {
			http.Error(w, "orderNumber query param required", http.StatusBadRequest)
			return
		}

		output, err := router.HandleGetOrderByNumber(r.Context(), controllers.GetOrderByNumberInput{OrderNumber: orderNumber})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"order":{"id":%d,"orderNumber":"%s","customerName":"%s"}}`, output.Order.ID, output.Order.OrderNumber, output.Order.CustomerName)
	}
}

// TMS endpoints
func handleTmsGetShipment(router *controllers.TmsRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		orderNumber := r.URL.Query().Get("orderNumber")
		if orderNumber == "" {
			http.Error(w, "orderNumber query param required", http.StatusBadRequest)
			return
		}

		output, err := router.HandleGetShipmentByOrderNumber(r.Context(), controllers.GetShipmentByOrderNumberInput{OrderNumber: orderNumber})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"shipment":{"id":%d,"shipmentNumber":"%s","status":"%s"}}`, output.Shipment.ID, output.Shipment.ShipmentNumber, output.Shipment.Status)
	}
}

// WMS endpoints
func handleWmsGetInventory(router *controllers.WmsRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		sku := r.URL.Query().Get("sku")
		if sku == "" {
			http.Error(w, "sku query param required", http.StatusBadRequest)
			return
		}

		output, err := router.HandleGetInventoryBySKU(r.Context(), controllers.GetInventoryBySKUInput{SKU: sku})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"inventory":{"id":%d,"sku":"%s","availableQuantity":%d}}`, output.Inventory.ID, output.Inventory.SKU, output.Inventory.AvailableQuantity)
	}
}

// SRM endpoints
func handleSrmGetSuppliers(router *controllers.SrmRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		category := r.URL.Query().Get("category")
		if category == "" {
			http.Error(w, "category query param required", http.StatusBadRequest)
			return
		}

		output, err := router.HandleGetSuppliersByCategory(r.Context(), controllers.GetSuppliersByCategoryInput{Category: category})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"suppliers":%v}`, output.Suppliers)
	}
}

func handleSrmGetPurchaseOrders(router *controllers.SrmRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		sku := r.URL.Query().Get("sku")
		if sku == "" {
			http.Error(w, "sku query param required", http.StatusBadRequest)
			return
		}

		output, err := router.HandleGetPurchaseOrdersBySKU(r.Context(), controllers.GetPurchaseOrdersBySKUInput{SKU: sku})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"orders":%v}`, output.Orders)
	}
}
