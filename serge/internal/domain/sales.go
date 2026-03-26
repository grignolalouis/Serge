package domain

import "time"

// ---------------------------------------------------------------------------
// Sales domain — customers and customer orders
// ---------------------------------------------------------------------------

type CustomerSegment string

const (
	SegmentRetail    CustomerSegment = "retail"
	SegmentWholesale CustomerSegment = "wholesale"
)

type Customer struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Company      string          `json:"company"`
	Segment      CustomerSegment `json:"segment"`
	Region       string          `json:"region"`
	Address      string          `json:"address"`
	DeliveryZone string          `json:"delivery_zone"`
}

type CustomerOrderStatus string

const (
	CustomerOrderPending    CustomerOrderStatus = "pending"
	CustomerOrderProcessing CustomerOrderStatus = "processing"
	CustomerOrderShipped    CustomerOrderStatus = "shipped"
	CustomerOrderDelivered  CustomerOrderStatus = "delivered"
	CustomerOrderCancelled  CustomerOrderStatus = "cancelled"
)

type Priority string

const (
	PriorityNormal Priority = "normal"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

// CustomerOrder is a sales order placed by a B2B customer.
type CustomerOrder struct {
	ID            string              `json:"id"`
	CustomerID    string              `json:"customer_id"`
	Status        CustomerOrderStatus `json:"status"`
	OrderDate     time.Time           `json:"order_date"`
	RequiredDate  time.Time           `json:"required_date"`
	ShippedDate   *time.Time          `json:"shipped_date,omitempty"`
	DeliveredDate *time.Time          `json:"delivered_date,omitempty"`
	Priority      Priority            `json:"priority"`
	Notes         string              `json:"notes,omitempty"`
	Lines         []CustomerOrderLine `json:"lines"`
}

// CustomerOrderLine references a Product — price comes from Product.UnitPrice.
type CustomerOrderLine struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (o CustomerOrder) IsOverdue() bool {
	return o.Status != CustomerOrderDelivered && o.Status != CustomerOrderCancelled && time.Now().After(o.RequiredDate)
}
