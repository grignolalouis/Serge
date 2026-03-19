package models

import "time"

// Shipment represents a shipment in the TMS system
type Shipment struct {
	ID                   uint
	ShipmentNumber       string
	OrderNumber          string
	Carrier              string
	Status               string
	EstimatedDeliveryDate *time.Time
}

// TableName specifies the schema and table name
func (Shipment) TableName() string {
	return "tms.shipments"
}
