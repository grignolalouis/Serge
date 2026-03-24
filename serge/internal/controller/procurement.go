package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

// ProcurementController exposes procurement operations for MCP tools (srm_*).
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

func (c *ProcurementController) GetPurchaseOrder(id string) (domain.PurchaseOrder, error) {
	return c.svc.GetPurchaseOrder(id)
}

func (c *ProcurementController) ListPurchaseOrdersBySupplier(supplierID string) []domain.PurchaseOrder {
	return c.svc.ListPurchaseOrdersBySupplier(supplierID)
}

func (c *ProcurementController) ListPurchaseOrdersByStatus(status string) []domain.PurchaseOrder {
	return c.svc.ListPurchaseOrdersByStatus(domain.PurchaseOrderStatus(status))
}

func (c *ProcurementController) ListOverduePurchaseOrders() []domain.PurchaseOrder {
	return c.svc.ListOverduePurchaseOrders()
}

func (c *ProcurementController) CompareSuppliers() []service.SupplierStats {
	return c.svc.CompareSuppliers()
}

func (c *ProcurementController) ListAllPurchaseOrders() []domain.PurchaseOrder {
	return c.svc.ListAllPurchaseOrders()
}
