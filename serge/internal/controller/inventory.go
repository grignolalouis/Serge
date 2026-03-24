package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

// InventoryController exposes inventory operations for MCP tools (wms_*).
type InventoryController struct {
	svc *service.InventoryService
}

func NewInventoryController(svc *service.InventoryService) *InventoryController {
	return &InventoryController{svc: svc}
}

func (c *InventoryController) GetProduct(id string) (domain.Product, error) {
	return c.svc.GetProduct(id)
}

func (c *InventoryController) GetProductBySKU(sku string) (domain.Product, error) {
	return c.svc.GetProductBySKU(sku)
}

func (c *InventoryController) ListProducts() []domain.Product {
	return c.svc.ListProducts()
}

func (c *InventoryController) CheckInventory(productID string) (service.InventoryStatus, error) {
	return c.svc.CheckInventory(productID)
}

func (c *InventoryController) ListLowStock() []service.InventoryStatus {
	return c.svc.ListLowStock()
}

func (c *InventoryController) GetStockMovements(productID string) []domain.StockMovement {
	return c.svc.GetStockMovements(productID)
}

func (c *InventoryController) ListAllInventory() []service.InventoryStatus {
	return c.svc.ListAllInventory()
}

func (c *InventoryController) GetWarehouse(id string) (domain.Warehouse, error) {
	return c.svc.GetWarehouse(id)
}
