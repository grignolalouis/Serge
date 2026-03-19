package models

// Inventory represents an inventory item in the WMS system
type Inventory struct {
	ID                uint
	SKU               string
	AvailableQuantity int
	ReservedQuantity  int
	Location          string
}

// TableName specifies the schema and table name
func (Inventory) TableName() string {
	return "wms.inventory"
}
