package controllers

import (
	"context"
	"serge/backend/internal/models"
	"serge/backend/internal/services"
)

// SrmRouter handles SRM-related requests
type SrmRouter struct {
	Service services.SrmService
}

// GetSuppliersByCategoryInput is the input for GetSuppliersByCategory
type GetSuppliersByCategoryInput struct {
	Category string
}

// GetSuppliersByCategoryOutput is the output for GetSuppliersByCategory
type GetSuppliersByCategoryOutput struct {
	Suppliers []models.Supplier
}

// HandleGetSuppliersByCategory handles getting suppliers by category
func (r *SrmRouter) HandleGetSuppliersByCategory(ctx context.Context, input GetSuppliersByCategoryInput) (GetSuppliersByCategoryOutput, error) {
	suppliers, err := r.Service.GetSuppliersByCategory(input.Category)
	if err != nil {
		return GetSuppliersByCategoryOutput{}, err
	}
	return GetSuppliersByCategoryOutput{Suppliers: suppliers}, nil
}

// GetPurchaseOrdersBySKUInput is the input for GetPurchaseOrdersBySKU
type GetPurchaseOrdersBySKUInput struct {
	SKU string
}

// GetPurchaseOrdersBySKUOutput is the output for GetPurchaseOrdersBySKU
type GetPurchaseOrdersBySKUOutput struct {
	Orders []models.PurchaseOrder
}

// HandleGetPurchaseOrdersBySKU handles getting purchase orders by SKU
func (r *SrmRouter) HandleGetPurchaseOrdersBySKU(ctx context.Context, input GetPurchaseOrdersBySKUInput) (GetPurchaseOrdersBySKUOutput, error) {
	orders, err := r.Service.GetPurchaseOrdersBySKU(input.SKU)
	if err != nil {
		return GetPurchaseOrdersBySKUOutput{}, err
	}
	return GetPurchaseOrdersBySKUOutput{Orders: orders}, nil
}
