package domain

import "time"

// ---------------------------------------------------------------------------
// Shipping domain — carriers, shipments, tracking
// ---------------------------------------------------------------------------

type Carrier struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	CostPerKg      float64 `json:"cost_per_kg"`
	AvgTransitDays int     `json:"avg_transit_days"`
}

type ShipmentStatus string

const (
	ShipmentPending   ShipmentStatus = "pending"
	ShipmentInTransit ShipmentStatus = "in_transit"
	ShipmentDelivered ShipmentStatus = "delivered"
	ShipmentException ShipmentStatus = "exception"
)

type Shipment struct {
	ID               string         `json:"id"`
	CustomerOrderID  string         `json:"customer_order_id"`
	CarrierID        string         `json:"carrier_id"`
	Status           ShipmentStatus `json:"status"`
	TrackingNumber   string         `json:"tracking_number"`
	Origin           string         `json:"origin"`
	Destination      string         `json:"destination"`
	WeightKg         float64        `json:"weight_kg"`
	ShippedAt        *time.Time     `json:"shipped_at,omitempty"`
	EstimatedArrival *time.Time     `json:"estimated_arrival,omitempty"`
	DeliveredAt      *time.Time     `json:"delivered_at,omitempty"`
	ExceptionReason  string         `json:"exception_reason,omitempty"`
}

func (s Shipment) IsLate() bool {
	if s.EstimatedArrival == nil {
		return false
	}
	if s.Status == ShipmentDelivered {
		return s.DeliveredAt != nil && s.DeliveredAt.After(*s.EstimatedArrival)
	}
	return time.Now().After(*s.EstimatedArrival)
}

type TrackingEvent struct {
	ShipmentID  string    `json:"shipment_id"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}
