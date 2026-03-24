package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type ProcurementRepository struct {
	suppliers      map[string]domain.Supplier
	purchaseOrders map[string]domain.PurchaseOrder
}

func NewProcurementRepository() *ProcurementRepository {
	return &ProcurementRepository{
		suppliers:      make(map[string]domain.Supplier),
		purchaseOrders: make(map[string]domain.PurchaseOrder),
	}
}

// --- Loaders (used by seed) ---

func (r *ProcurementRepository) LoadSupplier(s domain.Supplier) {
	r.suppliers[s.ID] = s
}

func (r *ProcurementRepository) LoadPurchaseOrder(po domain.PurchaseOrder) {
	r.purchaseOrders[po.ID] = po
}

// --- Queries ---

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

func (r *ProcurementRepository) GetPurchaseOrder(id string) (domain.PurchaseOrder, error) {
	po, ok := r.purchaseOrders[id]
	if !ok {
		return domain.PurchaseOrder{}, fmt.Errorf("purchase order %s not found", id)
	}
	return po, nil
}

func (r *ProcurementRepository) ListPurchaseOrders() []domain.PurchaseOrder {
	out := make([]domain.PurchaseOrder, 0, len(r.purchaseOrders))
	for _, po := range r.purchaseOrders {
		out = append(out, po)
	}
	return out
}

func (r *ProcurementRepository) ListPurchaseOrdersBySupplier(supplierID string) []domain.PurchaseOrder {
	var out []domain.PurchaseOrder
	for _, po := range r.purchaseOrders {
		if po.SupplierID == supplierID {
			out = append(out, po)
		}
	}
	return out
}

func (r *ProcurementRepository) ListPurchaseOrdersByStatus(status domain.PurchaseOrderStatus) []domain.PurchaseOrder {
	var out []domain.PurchaseOrder
	for _, po := range r.purchaseOrders {
		if po.Status == status {
			out = append(out, po)
		}
	}
	return out
}

func (r *ProcurementRepository) ListOverduePurchaseOrders() []domain.PurchaseOrder {
	var out []domain.PurchaseOrder
	for _, po := range r.purchaseOrders {
		if po.IsOverdue() {
			out = append(out, po)
		}
	}
	return out
}
