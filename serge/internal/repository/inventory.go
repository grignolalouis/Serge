package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type InventoryRepository struct {
	products       map[string]domain.Product
	warehouses     map[string]domain.Warehouse
	inventory      map[string]domain.InventoryRecord // key: productID
	stockMovements []domain.StockMovement
}

func NewInventoryRepository() *InventoryRepository {
	return &InventoryRepository{
		products:   make(map[string]domain.Product),
		warehouses: make(map[string]domain.Warehouse),
		inventory:  make(map[string]domain.InventoryRecord),
	}
}

// --- Loaders ---

func (r *InventoryRepository) LoadProduct(p domain.Product) {
	r.products[p.ID] = p
}

func (r *InventoryRepository) LoadWarehouse(w domain.Warehouse) {
	r.warehouses[w.ID] = w
}

func (r *InventoryRepository) LoadInventoryRecord(rec domain.InventoryRecord) {
	r.inventory[rec.ProductID] = rec
}

func (r *InventoryRepository) LoadStockMovement(m domain.StockMovement) {
	r.stockMovements = append(r.stockMovements, m)
}

// --- Queries ---

func (r *InventoryRepository) GetProduct(id string) (domain.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return domain.Product{}, fmt.Errorf("product %s not found", id)
	}
	return p, nil
}

func (r *InventoryRepository) GetProductBySKU(sku string) (domain.Product, error) {
	for _, p := range r.products {
		if p.SKU == sku {
			return p, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product with SKU %s not found", sku)
}

func (r *InventoryRepository) ListProducts() []domain.Product {
	out := make([]domain.Product, 0, len(r.products))
	for _, p := range r.products {
		out = append(out, p)
	}
	return out
}

func (r *InventoryRepository) GetInventory(productID string) (domain.InventoryRecord, error) {
	rec, ok := r.inventory[productID]
	if !ok {
		return domain.InventoryRecord{}, fmt.Errorf("inventory for product %s not found", productID)
	}
	return rec, nil
}

func (r *InventoryRepository) ListInventory() []domain.InventoryRecord {
	out := make([]domain.InventoryRecord, 0, len(r.inventory))
	for _, rec := range r.inventory {
		out = append(out, rec)
	}
	return out
}

func (r *InventoryRepository) ListLowStock() []domain.InventoryRecord {
	var out []domain.InventoryRecord
	for _, rec := range r.inventory {
		p, ok := r.products[rec.ProductID]
		if ok && rec.NeedsReorder(p.ReorderPoint) {
			out = append(out, rec)
		}
	}
	return out
}

func (r *InventoryRepository) ListStockMovements(productID string) []domain.StockMovement {
	var out []domain.StockMovement
	for _, m := range r.stockMovements {
		if m.ProductID == productID {
			out = append(out, m)
		}
	}
	return out
}

func (r *InventoryRepository) GetWarehouse(id string) (domain.Warehouse, error) {
	w, ok := r.warehouses[id]
	if !ok {
		return domain.Warehouse{}, fmt.Errorf("warehouse %s not found", id)
	}
	return w, nil
}
