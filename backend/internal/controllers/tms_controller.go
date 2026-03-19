package controllers

import (
	"context"
	"serge/backend/internal/models"
	"serge/backend/internal/services"
)

// TmsRouter handles TMS-related requests
type TmsRouter struct {
	Service services.TmsService
}

// GetShipmentByOrderNumberInput is the input for GetShipmentByOrderNumber
type GetShipmentByOrderNumberInput struct {
	OrderNumber string
}

// GetShipmentByOrderNumberOutput is the output for GetShipmentByOrderNumber
type GetShipmentByOrderNumberOutput struct {
	Shipment *models.Shipment
}

// HandleGetShipmentByOrderNumber handles getting a shipment by order number
func (r *TmsRouter) HandleGetShipmentByOrderNumber(ctx context.Context, input GetShipmentByOrderNumberInput) (GetShipmentByOrderNumberOutput, error) {
	shipment, err := r.Service.GetShipmentByOrderNumber(input.OrderNumber)
	if err != nil {
		return GetShipmentByOrderNumberOutput{}, err
	}
	return GetShipmentByOrderNumberOutput{Shipment: shipment}, nil
}
