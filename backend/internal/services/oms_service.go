package services

import (
	"fmt"
	"serge/backend/internal/models"
	"serge/backend/internal/repositories"
)

// OmsService defines methods for OMS business logic
type OmsService interface {
	GetOrderByNumber(orderNumber string) (*models.Order, error)
}

// omsService implements OmsService
type omsService struct {
	repo repositories.OmsRepository
}

// NewOmsService creates a new OMS service
func NewOmsService(repo repositories.OmsRepository) OmsService {
	return &omsService{repo: repo}
}

// GetOrderByNumber retrieves an order by number with error wrapping
func (s *omsService) GetOrderByNumber(orderNumber string) (*models.Order, error) {
	order, err := s.repo.GetOrderByNumber(orderNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by number: %w", err)
	}
	return order, nil
}
