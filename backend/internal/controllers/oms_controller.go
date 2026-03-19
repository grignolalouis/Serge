package controllers

import (
	"context"
	"serge/backend/internal/models"
	"serge/backend/internal/services"
)

// OmsRouter handles OMS-related requests
type OmsRouter struct {
	Service services.OmsService
}

// GetOrderByNumberInput is the input for GetOrderByNumber
type GetOrderByNumberInput struct {
	OrderNumber string
}

// GetOrderByNumberOutput is the output for GetOrderByNumber
type GetOrderByNumberOutput struct {
	Order *models.Order
}

// HandleGetOrderByNumber handles getting an order by number
func (r *OmsRouter) HandleGetOrderByNumber(ctx context.Context, input GetOrderByNumberInput) (GetOrderByNumberOutput, error) {
	order, err := r.Service.GetOrderByNumber(input.OrderNumber)
	if err != nil {
		return GetOrderByNumberOutput{}, err
	}
	return GetOrderByNumberOutput{Order: order}, nil
}
