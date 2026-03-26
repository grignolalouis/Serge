package controller

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/service"
)

type ShippingController struct {
	svc *service.ShippingService
}

func NewShippingController(svc *service.ShippingService) *ShippingController {
	return &ShippingController{svc: svc}
}

func (c *ShippingController) GetShipment(id string) (service.ShipmentDetail, error)              { return c.svc.GetShipment(id) }
func (c *ShippingController) TrackCustomerOrder(customerOrderID string) ([]service.ShipmentDetail, error) { return c.svc.TrackCustomerOrder(customerOrderID) }
func (c *ShippingController) ListExceptionShipments() []domain.Shipment                          { return c.svc.ListExceptionShipments() }
func (c *ShippingController) GetCarrier(id string) (domain.Carrier, error)                       { return c.svc.GetCarrier(id) }
func (c *ShippingController) ListShipmentsByStatus(status string) []domain.Shipment              { return c.svc.ListShipmentsByStatus(domain.ShipmentStatus(status)) }
func (c *ShippingController) ListCarriers() []domain.Carrier                                     { return c.svc.ListCarriers() }
