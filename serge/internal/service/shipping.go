package service

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

type ShipmentDetail struct {
	Shipment domain.Shipment        `json:"shipment"`
	Carrier  domain.Carrier         `json:"carrier"`
	Events   []domain.TrackingEvent `json:"tracking_events"`
}

type ShippingService struct {
	repo *repository.ShippingRepository
}

func NewShippingService(repo *repository.ShippingRepository) *ShippingService {
	return &ShippingService{repo: repo}
}

func (s *ShippingService) GetShipment(id string) (ShipmentDetail, error) {
	ship, err := s.repo.GetShipment(id)
	if err != nil {
		return ShipmentDetail{}, err
	}
	carrier, _ := s.repo.GetCarrier(ship.CarrierID)
	events := s.repo.GetTrackingEvents(id)
	return ShipmentDetail{Shipment: ship, Carrier: carrier, Events: events}, nil
}

func (s *ShippingService) TrackOrder(orderID string) (ShipmentDetail, error) {
	ship, err := s.repo.GetShipmentByOrder(orderID)
	if err != nil {
		return ShipmentDetail{}, err
	}
	carrier, _ := s.repo.GetCarrier(ship.CarrierID)
	events := s.repo.GetTrackingEvents(ship.ID)
	return ShipmentDetail{Shipment: ship, Carrier: carrier, Events: events}, nil
}

func (s *ShippingService) ListExceptionShipments() []domain.Shipment {
	return s.repo.ListExceptionShipments()
}

func (s *ShippingService) GetCarrier(id string) (domain.Carrier, error) {
	return s.repo.GetCarrier(id)
}
