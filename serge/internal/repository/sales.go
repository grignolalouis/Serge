package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type SalesRepository struct {
	customers map[string]domain.Customer
	orders    map[string]domain.Order
}

func NewSalesRepository() *SalesRepository {
	return &SalesRepository{
		customers: make(map[string]domain.Customer),
		orders:    make(map[string]domain.Order),
	}
}

// --- Loaders ---

func (r *SalesRepository) LoadCustomer(c domain.Customer) {
	r.customers[c.ID] = c
}

func (r *SalesRepository) LoadOrder(o domain.Order) {
	r.orders[o.ID] = o
}

// --- Queries ---

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

func (r *SalesRepository) GetOrder(id string) (domain.Order, error) {
	o, ok := r.orders[id]
	if !ok {
		return domain.Order{}, fmt.Errorf("order %s not found", id)
	}
	return o, nil
}

func (r *SalesRepository) ListOrders() []domain.Order {
	out := make([]domain.Order, 0, len(r.orders))
	for _, o := range r.orders {
		out = append(out, o)
	}
	return out
}

func (r *SalesRepository) ListOrdersByCustomer(customerID string) []domain.Order {
	var out []domain.Order
	for _, o := range r.orders {
		if o.CustomerID == customerID {
			out = append(out, o)
		}
	}
	return out
}

func (r *SalesRepository) ListOrdersByStatus(status domain.OrderStatus) []domain.Order {
	var out []domain.Order
	for _, o := range r.orders {
		if o.Status == status {
			out = append(out, o)
		}
	}
	return out
}

func (r *SalesRepository) ListOverdueOrders() []domain.Order {
	var out []domain.Order
	for _, o := range r.orders {
		if o.IsOverdue() {
			out = append(out, o)
		}
	}
	return out
}
