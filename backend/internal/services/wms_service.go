package services

import (
	"fmt"
	"serge/backend/internal/models"
	"serge/backend/internal/repositories"
)

// WmsService defines methods for WMS business logic
type WmsService interface {
	GetInventoryBySKU(sku string) (*models.Inventory, error)
}

// wmsService implements WmsService
type wmsService struct {
	repo repositories.WmsRepository
}

// NewWmsService creates a new WMS service
func NewWmsService(repo repositories.WmsRepository) WmsService {
	return &wmsService{repo: repo}
}

// GetInventoryBySKU retrieves inventory by SKU with error wrapping
func (s *wmsService) GetInventoryBySKU(sku string) (*models.Inventory, error) {
	inventory, err := s.repo.GetInventoryBySKU(sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory by SKU: %w", err)
	}
	return inventory, nil
}
