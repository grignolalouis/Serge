package seed

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

func LoadAll(dataDir string, procurement *repository.ProcurementRepository, inventory *repository.InventoryRepository, sales *repository.SalesRepository, shipping *repository.ShippingRepository) error {
	// Procurement
	var suppliers []domain.Supplier
	if err := loadJSON(filepath.Join(dataDir, "suppliers.json"), &suppliers); err != nil {
		return fmt.Errorf("suppliers: %w", err)
	}
	for _, s := range suppliers {
		procurement.LoadSupplier(s)
	}

	var pos []domain.PurchaseOrder
	if err := loadJSON(filepath.Join(dataDir, "purchase_orders.json"), &pos); err != nil {
		return fmt.Errorf("purchase_orders: %w", err)
	}
	for _, po := range pos {
		procurement.LoadPurchaseOrder(po)
	}

	// Inventory
	var products []domain.Product
	if err := loadJSON(filepath.Join(dataDir, "products.json"), &products); err != nil {
		return fmt.Errorf("products: %w", err)
	}
	for _, p := range products {
		inventory.LoadProduct(p)
	}

	var warehouses []domain.Warehouse
	if err := loadJSON(filepath.Join(dataDir, "warehouses.json"), &warehouses); err != nil {
		return fmt.Errorf("warehouses: %w", err)
	}
	for _, w := range warehouses {
		inventory.LoadWarehouse(w)
	}

	var records []domain.InventoryRecord
	if err := loadJSON(filepath.Join(dataDir, "inventory.json"), &records); err != nil {
		return fmt.Errorf("inventory: %w", err)
	}
	for _, r := range records {
		inventory.LoadInventoryRecord(r)
	}

	var movements []domain.StockMovement
	if err := loadJSON(filepath.Join(dataDir, "stock_movements.json"), &movements); err != nil {
		return fmt.Errorf("stock_movements: %w", err)
	}
	for _, m := range movements {
		inventory.LoadStockMovement(m)
	}

	// Sales
	var customers []domain.Customer
	if err := loadJSON(filepath.Join(dataDir, "customers.json"), &customers); err != nil {
		return fmt.Errorf("customers: %w", err)
	}
	for _, c := range customers {
		sales.LoadCustomer(c)
	}

	var orders []domain.Order
	if err := loadJSON(filepath.Join(dataDir, "orders.json"), &orders); err != nil {
		return fmt.Errorf("orders: %w", err)
	}
	for _, o := range orders {
		sales.LoadOrder(o)
	}

	// Shipping
	var carriers []domain.Carrier
	if err := loadJSON(filepath.Join(dataDir, "carriers.json"), &carriers); err != nil {
		return fmt.Errorf("carriers: %w", err)
	}
	for _, c := range carriers {
		shipping.LoadCarrier(c)
	}

	var shipments []domain.Shipment
	if err := loadJSON(filepath.Join(dataDir, "shipments.json"), &shipments); err != nil {
		return fmt.Errorf("shipments: %w", err)
	}
	for _, s := range shipments {
		shipping.LoadShipment(s)
	}

	var events []domain.TrackingEvent
	if err := loadJSON(filepath.Join(dataDir, "tracking_events.json"), &events); err != nil {
		return fmt.Errorf("tracking_events: %w", err)
	}
	for _, e := range events {
		shipping.LoadTrackingEvent(e)
	}

	return nil
}

func loadJSON(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
