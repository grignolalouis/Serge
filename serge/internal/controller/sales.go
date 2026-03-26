package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

type SalesController struct {
	svc *service.SalesService
}

func NewSalesController(svc *service.SalesService) *SalesController {
	return &SalesController{svc: svc}
}

func (c *SalesController) GetCustomerOrder(id string) (domain.CustomerOrder, error) { return c.svc.GetCustomerOrder(id) }
func (c *SalesController) GetCustomer(id string) (domain.Customer, error)           { return c.svc.GetCustomer(id) }
func (c *SalesController) ListCustomerOrdersByCustomer(customerID string) []domain.CustomerOrder { return c.svc.ListCustomerOrdersByCustomer(customerID) }
func (c *SalesController) ListCustomerOrdersByStatus(status string) []domain.CustomerOrder { return c.svc.ListCustomerOrdersByStatus(domain.CustomerOrderStatus(status)) }
func (c *SalesController) ListOverdueCustomerOrders() []domain.CustomerOrder       { return c.svc.ListOverdueCustomerOrders() }
func (c *SalesController) ListAllCustomerOrders() []domain.CustomerOrder           { return c.svc.ListAllCustomerOrders() }
func (c *SalesController) ListCustomers() []domain.Customer                        { return c.svc.ListCustomers() }
func (c *SalesController) ListCustomerOrdersByProduct(productID string) []domain.CustomerOrder { return c.svc.ListCustomerOrdersByProduct(productID) }
