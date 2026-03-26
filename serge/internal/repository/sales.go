package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type SalesRepository struct {
	customers      map[string]domain.Customer
	customerOrders map[string]domain.CustomerOrder
}

func NewSalesRepository() *SalesRepository {
	return &SalesRepository{
		customers:      make(map[string]domain.Customer),
		customerOrders: make(map[string]domain.CustomerOrder),
	}
}

func (r *SalesRepository) LoadCustomer(c domain.Customer)           { r.customers[c.ID] = c }
func (r *SalesRepository) LoadCustomerOrder(o domain.CustomerOrder) { r.customerOrders[o.ID] = o }

func (r *SalesRepository) GetCustomer(id string) (domain.Customer, error) {
	c, ok := r.customers[id]
	if !ok {
		return domain.Customer{}, fmt.Errorf("customer %s not found", id)
	}
	return c, nil
}

func (r *SalesRepository) ListCustomers() []domain.Customer {
	out := make([]domain.Customer, 0, len(r.customers))
	for _, c := range r.customers {
		out = append(out, c)
	}
	return out
}

func (r *SalesRepository) GetCustomerOrder(id string) (domain.CustomerOrder, error) {
	o, ok := r.customerOrders[id]
	if !ok {
		return domain.CustomerOrder{}, fmt.Errorf("customer order %s not found", id)
	}
	return o, nil
}

func (r *SalesRepository) ListCustomerOrders() []domain.CustomerOrder {
	out := make([]domain.CustomerOrder, 0, len(r.customerOrders))
	for _, o := range r.customerOrders {
		out = append(out, o)
	}
	return out
}

func (r *SalesRepository) ListCustomerOrdersByCustomer(customerID string) []domain.CustomerOrder {
	var out []domain.CustomerOrder
	for _, o := range r.customerOrders {
		if o.CustomerID == customerID {
			out = append(out, o)
		}
	}
	return out
}

func (r *SalesRepository) ListCustomerOrdersByStatus(status domain.CustomerOrderStatus) []domain.CustomerOrder {
	var out []domain.CustomerOrder
	for _, o := range r.customerOrders {
		if o.Status == status {
			out = append(out, o)
		}
	}
	return out
}

func (r *SalesRepository) ListOverdueCustomerOrders() []domain.CustomerOrder {
	var out []domain.CustomerOrder
	for _, o := range r.customerOrders {
		if o.IsOverdue() {
			out = append(out, o)
		}
	}
	return out
}

func (r *SalesRepository) ListCustomerOrdersByProduct(productID string) []domain.CustomerOrder {
	var out []domain.CustomerOrder
	for _, o := range r.customerOrders {
		for _, l := range o.Lines {
			if l.ProductID == productID {
				out = append(out, o)
				break
			}
		}
	}
	return out
}
