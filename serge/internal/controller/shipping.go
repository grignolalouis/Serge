package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

// ShippingController exposes shipping operations for MCP tools (tms_*).
type ShippingController struct {
	svc *service.ShippingService
}

func NewShippingController(svc *service.ShippingService) *ShippingController {
	return &ShippingController{svc: svc}
}

func (c *ShippingController) GetShipment(id string) (service.ShipmentDetail, error) {
	return c.svc.GetShipment(id)
}

func (c *ShippingController) TrackOrder(orderID string) (service.ShipmentDetail, error) {
	return c.svc.TrackOrder(orderID)
}

func (c *ShippingController) ListExceptionShipments() []domain.Shipment {
	return c.svc.ListExceptionShipments()
}

func (c *ShippingController) GetCarrier(id string) (domain.Carrier, error) {
	return c.svc.GetCarrier(id)
}
