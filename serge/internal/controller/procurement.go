package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

type ProcurementController struct {
	svc *service.ProcurementService
}

func NewProcurementController(svc *service.ProcurementService) *ProcurementController {
	return &ProcurementController{svc: svc}
}

func (c *ProcurementController) GetSupplier(id string) (domain.Supplier, error) {
	return c.svc.GetSupplier(id)
}

func (c *ProcurementController) ListSuppliers() []domain.Supplier {
	return c.svc.ListSuppliers()
}

func (c *ProcurementController) GetSupplierProduct(id string) (domain.SupplierProduct, error) {
	return c.svc.GetSupplierProduct(id)
}

func (c *ProcurementController) ListSupplierProductsBySupplier(supplierID string) []domain.SupplierProduct {
	return c.svc.ListSupplierProductsBySupplier(supplierID)
}

func (c *ProcurementController) ListSupplierProductsByProduct(productID string) []domain.SupplierProduct {
	return c.svc.ListSupplierProductsByProduct(productID)
}

func (c *ProcurementController) GetSupplierOrder(id string) (domain.SupplierOrder, error) {
	return c.svc.GetSupplierOrder(id)
}

func (c *ProcurementController) ListAllSupplierOrders() []domain.SupplierOrder {
	return c.svc.ListAllSupplierOrders()
}

func (c *ProcurementController) ListSupplierOrdersBySupplier(supplierID string) []domain.SupplierOrder {
	return c.svc.ListSupplierOrdersBySupplier(supplierID)
}

func (c *ProcurementController) ListSupplierOrdersByStatus(status string) []domain.SupplierOrder {
	return c.svc.ListSupplierOrdersByStatus(domain.SupplierOrderStatus(status))
}

func (c *ProcurementController) ListOverdueSupplierOrders() []domain.SupplierOrder {
	return c.svc.ListOverdueSupplierOrders()
}

func (c *ProcurementController) CompareSuppliers() []service.SupplierStats {
	return c.svc.CompareSuppliers()
}
