package domain

import "time"

type CustomerSegment string

const (
	SegmentRetail    CustomerSegment = "retail"
	SegmentWholesale CustomerSegment = "wholesale"
)

type Customer struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Company string          `json:"company"`
	Segment CustomerSegment `json:"segment"`
	Region  string          `json:"region"`
}

type OrderStatus string

const (
	OrderPending    OrderStatus = "pending"
	OrderProcessing OrderStatus = "processing"
	OrderShipped    OrderStatus = "shipped"
	OrderDelivered  OrderStatus = "delivered"
	OrderCancelled  OrderStatus = "cancelled"
)

type Priority string

const (
	PriorityNormal Priority = "normal"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

type Order struct {
	ID            string      `json:"id"`
	CustomerID    string      `json:"customer_id"`
	Status        OrderStatus `json:"status"`
	OrderDate     time.Time   `json:"order_date"`
	RequiredDate  time.Time   `json:"required_date"`
	ShippedDate   *time.Time  `json:"shipped_date,omitempty"`
	DeliveredDate *time.Time  `json:"delivered_date,omitempty"`
	Priority      Priority    `json:"priority"`
	Lines         []OrderLine `json:"lines"`
}

type OrderLine struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

func (o Order) TotalAmount() float64 {
	var total float64
	for _, l := range o.Lines {
		total += float64(l.Quantity) * l.UnitPrice
	}
	return total
}

func (o Order) IsOverdue() bool {
	return o.Status != OrderDelivered && o.Status != OrderCancelled && time.Now().After(o.RequiredDate)
}
