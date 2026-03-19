package controllers

import (
	"context"
	"serge/backend/internal/models"
	"serge/backend/internal/services"
)

// WmsRouter handles WMS-related requests
type WmsRouter struct {
	Service services.WmsService
}

// GetInventoryBySKUInput is the input for GetInventoryBySKU
type GetInventoryBySKUInput struct {
	SKU string
}

// GetInventoryBySKUOutput is the output for GetInventoryBySKU
type GetInventoryBySKUOutput struct {
	Inventory *models.Inventory
}

// HandleGetInventoryBySKU handles getting inventory by SKU
func (r *WmsRouter) HandleGetInventoryBySKU(ctx context.Context, input GetInventoryBySKUInput) (GetInventoryBySKUOutput, error) {
	inventory, err := r.Service.GetInventoryBySKU(input.SKU)
	if err != nil {
		return GetInventoryBySKUOutput{}, err
	}
	return GetInventoryBySKUOutput{Inventory: inventory}, nil
}
