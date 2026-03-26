package seed

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/serge-music/serge/internal/repository"
)

func LoadAll(dataDir string, procurement *repository.ProcurementRepository, inventory *repository.InventoryRepository, sales *repository.SalesRepository, shipping *repository.ShippingRepository) error {
	if err := loadEntities(dataDir, "suppliers.json", procurement.LoadSupplier); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "supplier_products.json", procurement.LoadSupplierProduct); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "supplier_orders.json", procurement.LoadSupplierOrder); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "products.json", inventory.LoadProduct); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "warehouses.json", inventory.LoadWarehouse); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "inventory.json", inventory.LoadInventoryRecord); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "stock_movements.json", inventory.LoadStockMovement); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "customers.json", sales.LoadCustomer); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "customer_orders.json", sales.LoadCustomerOrder); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "carriers.json", shipping.LoadCarrier); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "shipments.json", shipping.LoadShipment); err != nil {
		return err
	}
	if err := loadEntities(dataDir, "tracking_events.json", shipping.LoadTrackingEvent); err != nil {
		return err
	}
	return nil
}

func loadEntities[T any](dataDir, filename string, load func(T)) error {
	path := filepath.Join(dataDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%s: %w", filename, err)
	}
	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("%s: %w", filename, err)
	}
	for _, item := range items {
		load(item)
	}
	return nil
}
