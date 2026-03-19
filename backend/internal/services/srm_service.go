package services

import (
	"fmt"
	"serge/backend/internal/models"
	"serge/backend/internal/repositories"
)

// SrmService defines methods for SRM business logic
type SrmService interface {
	GetSuppliersByCategory(category string) ([]models.Supplier, error)
	GetPurchaseOrdersBySKU(sku string) ([]models.PurchaseOrder, error)
}

// srmService implements SrmService
type srmService struct {
	repo repositories.SrmRepository
}

// NewSrmService creates a new SRM service
func NewSrmService(repo repositories.SrmRepository) SrmService {
	return &srmService{repo: repo}
}

// GetSuppliersByCategory retrieves suppliers by category with error wrapping
func (s *srmService) GetSuppliersByCategory(category string) ([]models.Supplier, error) {
	suppliers, err := s.repo.GetSuppliersByCategory(category)
	if err != nil {
		return nil, fmt.Errorf("failed to get suppliers by category: %w", err)
	}
	return suppliers, nil
}

// GetPurchaseOrdersBySKU retrieves purchase orders by SKU with error wrapping
func (s *srmService) GetPurchaseOrdersBySKU(sku string) ([]models.PurchaseOrder, error) {
	orders, err := s.repo.GetPurchaseOrdersBySKU(sku)
	if err != nil {
		return nil, fmt.Errorf("failed to get purchase orders by SKU: %w", err)
	}
	return orders, nil
}
