package models

import "time"

// Supplier represents a supplier in the SRM system
type Supplier struct {
	ID              uint
	Name            string
	Category        string
	OnTimeRatePct   float64
	AvgLeadTimeDays int
	DefectRatePct   float64
}

// TableName specifies the schema and table name
func (Supplier) TableName() string {
	return "srm.suppliers"
}

// PurchaseOrder represents a purchase order in the SRM system
type PurchaseOrder struct {
	ID                  uint
	PONumber            string
	SupplierID          uint
	SKU                 string
	Quantity            int
	Status              string
	ExpectedDeliveryDate *time.Time
}

// TableName specifies the schema and table name
func (PurchaseOrder) TableName() string {
	return "srm.purchase_orders"
}
