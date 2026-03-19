package models

import "time"

// Order represents an order in the OMS system
type Order struct {
	ID            uint
	OrderNumber   string
	CustomerName  string
	Status        string
	PlacementDate time.Time
	RequiredDate  time.Time
	Lines         []OrderLine
}

// TableName specifies the schema and table name
func (Order) TableName() string {
	return "oms.orders"
}

// OrderLine represents a line item in an order
type OrderLine struct {
	ID       uint
	OrderID  uint
	SKU      string
	Quantity int
}

// TableName specifies the schema and table name
func (OrderLine) TableName() string {
	return "oms.order_lines"
}
