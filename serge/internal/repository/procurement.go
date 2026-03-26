package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type ProcurementRepository struct {
	suppliers        map[string]domain.Supplier
	supplierProducts map[string]domain.SupplierProduct
	supplierOrders   map[string]domain.SupplierOrder
}

func NewProcurementRepository() *ProcurementRepository {
	return &ProcurementRepository{
		suppliers:        make(map[string]domain.Supplier),
		supplierProducts: make(map[string]domain.SupplierProduct),
		supplierOrders:   make(map[string]domain.SupplierOrder),
	}
}

// --- Loaders ---

func (r *ProcurementRepository) LoadSupplier(s domain.Supplier)               { r.suppliers[s.ID] = s }
func (r *ProcurementRepository) LoadSupplierProduct(sp domain.SupplierProduct) { r.supplierProducts[sp.ID] = sp }
func (r *ProcurementRepository) LoadSupplierOrder(so domain.SupplierOrder)     { r.supplierOrders[so.ID] = so }

// --- Supplier queries ---

func (r *ProcurementRepository) GetSupplier(id string) (domain.Supplier, error) {
	s, ok := r.suppliers[id]
	if !ok {
		return domain.Supplier{}, fmt.Errorf("supplier %s not found", id)
	}
	return s, nil
}

func (r *ProcurementRepository) ListSuppliers() []domain.Supplier {
	out := make([]domain.Supplier, 0, len(r.suppliers))
	for _, s := range r.suppliers {
		out = append(out, s)
	}
	return out
}

// --- SupplierProduct queries ---

func (r *ProcurementRepository) GetSupplierProduct(id string) (domain.SupplierProduct, error) {
	sp, ok := r.supplierProducts[id]
	if !ok {
		return domain.SupplierProduct{}, fmt.Errorf("supplier product %s not found", id)
	}
	return sp, nil
}

func (r *ProcurementRepository) ListSupplierProducts() []domain.SupplierProduct {
	out := make([]domain.SupplierProduct, 0, len(r.supplierProducts))
	for _, sp := range r.supplierProducts {
		out = append(out, sp)
	}
	return out
}

func (r *ProcurementRepository) ListSupplierProductsBySupplier(supplierID string) []domain.SupplierProduct {
	var out []domain.SupplierProduct
	for _, sp := range r.supplierProducts {
		if sp.SupplierID == supplierID {
			out = append(out, sp)
		}
	}
	return out
}

func (r *ProcurementRepository) ListSupplierProductsByProduct(productID string) []domain.SupplierProduct {
	var out []domain.SupplierProduct
	for _, sp := range r.supplierProducts {
		if sp.ProductID == productID {
			out = append(out, sp)
		}
	}
	return out
}

// --- SupplierOrder queries ---

func (r *ProcurementRepository) GetSupplierOrder(id string) (domain.SupplierOrder, error) {
	so, ok := r.supplierOrders[id]
	if !ok {
		return domain.SupplierOrder{}, fmt.Errorf("supplier order %s not found", id)
	}
	return so, nil
}

func (r *ProcurementRepository) ListSupplierOrders() []domain.SupplierOrder {
	out := make([]domain.SupplierOrder, 0, len(r.supplierOrders))
	for _, so := range r.supplierOrders {
		out = append(out, so)
	}
	return out
}

func (r *ProcurementRepository) ListSupplierOrdersBySupplier(supplierID string) []domain.SupplierOrder {
	var out []domain.SupplierOrder
	for _, so := range r.supplierOrders {
		if so.SupplierID == supplierID {
			out = append(out, so)
		}
	}
	return out
}

func (r *ProcurementRepository) ListSupplierOrdersByStatus(status domain.SupplierOrderStatus) []domain.SupplierOrder {
	var out []domain.SupplierOrder
	for _, so := range r.supplierOrders {
		if so.Status == status {
			out = append(out, so)
		}
	}
	return out
}

func (r *ProcurementRepository) ListOverdueSupplierOrders() []domain.SupplierOrder {
	var out []domain.SupplierOrder
	for _, so := range r.supplierOrders {
		if so.IsOverdue() {
			out = append(out, so)
		}
	}
	return out
}
