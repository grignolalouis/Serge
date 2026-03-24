package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

// SalesController exposes sales operations for MCP tools (oms_*).
type SalesController struct {
	svc *service.SalesService
}

func NewSalesController(svc *service.SalesService) *SalesController {
	return &SalesController{svc: svc}
}

func (c *SalesController) GetOrder(id string) (domain.Order, error) {
	return c.svc.GetOrder(id)
}

func (c *SalesController) GetCustomer(id string) (domain.Customer, error) {
	return c.svc.GetCustomer(id)
}

func (c *SalesController) ListOrdersByCustomer(customerID string) []domain.Order {
	return c.svc.ListOrdersByCustomer(customerID)
}

func (c *SalesController) ListOrdersByStatus(status string) []domain.Order {
	return c.svc.ListOrdersByStatus(domain.OrderStatus(status))
}

func (c *SalesController) ListOverdueOrders() []domain.Order {
	return c.svc.ListOverdueOrders()
}

func (c *SalesController) ListAllOrders() []domain.Order {
	return c.svc.ListAllOrders()
}

func (c *SalesController) ListCustomers() []domain.Customer {
	return c.svc.ListCustomers()
}

func (c *SalesController) ListOrdersByProduct(productID string) []domain.Order {
	return c.svc.ListOrdersByProduct(productID)
}
