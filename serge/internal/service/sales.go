package service

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

type SalesService struct {
	repo *repository.SalesRepository
}

func NewSalesService(repo *repository.SalesRepository) *SalesService {
	return &SalesService{repo: repo}
}

func (s *SalesService) GetOrder(id string) (domain.Order, error) {
	return s.repo.GetOrder(id)
}

func (s *SalesService) GetCustomer(id string) (domain.Customer, error) {
	return s.repo.GetCustomer(id)
}

func (s *SalesService) ListOrdersByCustomer(customerID string) []domain.Order {
	return s.repo.ListOrdersByCustomer(customerID)
}

func (s *SalesService) ListOrdersByStatus(status domain.OrderStatus) []domain.Order {
	return s.repo.ListOrdersByStatus(status)
}

func (s *SalesService) ListOverdueOrders() []domain.Order {
	return s.repo.ListOverdueOrders()
}

func (s *SalesService) ListAllOrders() []domain.Order {
	return s.repo.ListOrders()
}

func (s *SalesService) ListCustomers() []domain.Customer {
	return s.repo.ListCustomers()
}

func (s *SalesService) ListOrdersByProduct(productID string) []domain.Order {
	return s.repo.ListOrdersByProduct(productID)
}
