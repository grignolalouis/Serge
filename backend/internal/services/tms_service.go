package services

import (
	"fmt"
	"serge/backend/internal/models"
	"serge/backend/internal/repositories"
)

// TmsService defines methods for TMS business logic
type TmsService interface {
	GetShipmentByOrderNumber(orderNumber string) (*models.Shipment, error)
}

// tmsService implements TmsService
type tmsService struct {
	repo repositories.TmsRepository
}

// NewTmsService creates a new TMS service
func NewTmsService(repo repositories.TmsRepository) TmsService {
	return &tmsService{repo: repo}
}

// GetShipmentByOrderNumber retrieves a shipment by order number with error wrapping
func (s *tmsService) GetShipmentByOrderNumber(orderNumber string) (*models.Shipment, error) {
	shipment, err := s.repo.GetShipmentByOrderNumber(orderNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get shipment by order number: %w", err)
	}
	return shipment, nil
}
