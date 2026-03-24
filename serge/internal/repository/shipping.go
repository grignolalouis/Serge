package repository

import (
	"fmt"

	"github.com/serge-music/serge/internal/domain"
)

type ShippingRepository struct {
	carriers       map[string]domain.Carrier
	shipments      map[string]domain.Shipment
	trackingEvents []domain.TrackingEvent
}

func NewShippingRepository() *ShippingRepository {
	return &ShippingRepository{
		carriers:  make(map[string]domain.Carrier),
		shipments: make(map[string]domain.Shipment),
	}
}

// --- Loaders ---

func (r *ShippingRepository) LoadCarrier(c domain.Carrier) {
	r.carriers[c.ID] = c
}

func (r *ShippingRepository) LoadShipment(s domain.Shipment) {
	r.shipments[s.ID] = s
}

func (r *ShippingRepository) LoadTrackingEvent(e domain.TrackingEvent) {
	r.trackingEvents = append(r.trackingEvents, e)
}

// --- Queries ---

func (r *ShippingRepository) GetCarrier(id string) (domain.Carrier, error) {
	c, ok := r.carriers[id]
	if !ok {
		return domain.Carrier{}, fmt.Errorf("carrier %s not found", id)
	}
	return c, nil
}

func (r *ShippingRepository) ListCarriers() []domain.Carrier {
	out := make([]domain.Carrier, 0, len(r.carriers))
	for _, c := range r.carriers {
		out = append(out, c)
	}
	return out
}

func (r *ShippingRepository) GetShipment(id string) (domain.Shipment, error) {
	s, ok := r.shipments[id]
	if !ok {
		return domain.Shipment{}, fmt.Errorf("shipment %s not found", id)
	}
	return s, nil
}

func (r *ShippingRepository) GetShipmentByOrder(orderID string) (domain.Shipment, error) {
	for _, s := range r.shipments {
		if s.OrderID == orderID {
			return s, nil
		}
	}
	return domain.Shipment{}, fmt.Errorf("no shipment found for order %s", orderID)
}

func (r *ShippingRepository) ListShipments() []domain.Shipment {
	out := make([]domain.Shipment, 0, len(r.shipments))
	for _, s := range r.shipments {
		out = append(out, s)
	}
	return out
}

func (r *ShippingRepository) ListShipmentsByStatus(status domain.ShipmentStatus) []domain.Shipment {
	var out []domain.Shipment
	for _, s := range r.shipments {
		if s.Status == status {
			out = append(out, s)
		}
	}
	return out
}

func (r *ShippingRepository) ListExceptionShipments() []domain.Shipment {
	return r.ListShipmentsByStatus(domain.ShipmentException)
}

func (r *ShippingRepository) GetTrackingEvents(shipmentID string) []domain.TrackingEvent {
	var out []domain.TrackingEvent
	for _, e := range r.trackingEvents {
		if e.ShipmentID == shipmentID {
			out = append(out, e)
		}
	}
	return out
}
