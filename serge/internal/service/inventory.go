package service

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

type InventoryStatus struct {
	Product      domain.Product         `json:"product"`
	Record       domain.InventoryRecord `json:"inventory"`
	Available    int                    `json:"available"`
	NeedsReorder bool                  `json:"needs_reorder"`
}

type InventoryService struct {
	repo *repository.InventoryRepository
}

func NewInventoryService(repo *repository.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) GetProduct(id string) (domain.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *InventoryService) GetProductBySKU(sku string) (domain.Product, error) {
	return s.repo.GetProductBySKU(sku)
}

func (s *InventoryService) ListProducts() []domain.Product {
	return s.repo.ListProducts()
}

func (s *InventoryService) CheckInventory(productID string) (InventoryStatus, error) {
	p, err := s.repo.GetProduct(productID)
	if err != nil {
		return InventoryStatus{}, err
	}
	rec, err := s.repo.GetInventory(productID)
	if err != nil {
		return InventoryStatus{}, err
	}
	return InventoryStatus{
		Product:      p,
		Record:       rec,
		Available:    rec.Available(),
		NeedsReorder: rec.NeedsReorder(p.ReorderPoint),
	}, nil
}

func (s *InventoryService) ListLowStock() []InventoryStatus {
	records := s.repo.ListLowStock()
	var out []InventoryStatus
	for _, rec := range records {
		p, err := s.repo.GetProduct(rec.ProductID)
		if err != nil {
			continue
		}
		out = append(out, InventoryStatus{
			Product:      p,
			Record:       rec,
			Available:    rec.Available(),
			NeedsReorder: true,
		})
	}
	return out
}

func (s *InventoryService) GetStockMovements(productID string) []domain.StockMovement {
	return s.repo.ListStockMovements(productID)
}

func (s *InventoryService) ListAllInventory() []InventoryStatus {
	records := s.repo.ListInventory()
	var out []InventoryStatus
	for _, rec := range records {
		p, err := s.repo.GetProduct(rec.ProductID)
		if err != nil {
			continue
		}
		out = append(out, InventoryStatus{
			Product:      p,
			Record:       rec,
			Available:    rec.Available(),
			NeedsReorder: rec.NeedsReorder(p.ReorderPoint),
		})
	}
	return out
}

func (s *InventoryService) GetWarehouse(id string) (domain.Warehouse, error) {
	return s.repo.GetWarehouse(id)
}
