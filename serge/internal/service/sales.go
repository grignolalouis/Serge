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

func (s *SalesService) GetCustomerOrder(id string) (domain.CustomerOrder, error) {
	return s.repo.GetCustomerOrder(id)
}

func (s *SalesService) GetCustomer(id string) (domain.Customer, error) {
	return s.repo.GetCustomer(id)
}

func (s *SalesService) ListCustomerOrdersByCustomer(customerID string) []domain.CustomerOrder {
	return s.repo.ListCustomerOrdersByCustomer(customerID)
}

func (s *SalesService) ListCustomerOrdersByStatus(status domain.CustomerOrderStatus) []domain.CustomerOrder {
	return s.repo.ListCustomerOrdersByStatus(status)
}

func (s *SalesService) ListOverdueCustomerOrders() []domain.CustomerOrder {
	return s.repo.ListOverdueCustomerOrders()
}

func (s *SalesService) ListAllCustomerOrders() []domain.CustomerOrder {
	return s.repo.ListCustomerOrders()
}

func (s *SalesService) ListCustomers() []domain.Customer {
	return s.repo.ListCustomers()
}

func (s *SalesService) ListCustomerOrdersByProduct(productID string) []domain.CustomerOrder {
	return s.repo.ListCustomerOrdersByProduct(productID)
}
