package domain

import "time"

// ---------------------------------------------------------------------------
// Procurement domain — suppliers, their products, and supplier orders
// ---------------------------------------------------------------------------

type Supplier struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Region            string   `json:"region"`
	Contact           string   `json:"contact"`
	LeadTimeDays      int      `json:"lead_time_days"`
	Rating            float64  `json:"rating"`
	ProductCategories []string `json:"product_categories"`
	PaymentTerms      string   `json:"payment_terms"`
}

// SupplierProduct is an item a supplier offers for sale.
// UnitCost is the supplier's price — it is NOT duplicated on order lines.
type SupplierProduct struct {
	ID         string  `json:"id"`
	SupplierID string  `json:"supplier_id"`
	ProductID  string  `json:"product_id"` // our Product this feeds into
	Name       string  `json:"name"`       // supplier's name for this item
	UnitCost   float64 `json:"unit_cost"`  // supplier's selling price
}

type SupplierOrderStatus string

const (
	SupplierOrderPending   SupplierOrderStatus = "pending"
	SupplierOrderConfirmed SupplierOrderStatus = "confirmed"
	SupplierOrderShipped   SupplierOrderStatus = "shipped"
	SupplierOrderReceived  SupplierOrderStatus = "received"
	SupplierOrderCancelled SupplierOrderStatus = "cancelled"
)

// SupplierOrder is a purchase we place with a supplier to restock inventory.
type SupplierOrder struct {
	ID               string              `json:"id"`
	SupplierID       string              `json:"supplier_id"`
	Status           SupplierOrderStatus `json:"status"`
	OrderDate        time.Time           `json:"order_date"`
	ExpectedDelivery time.Time           `json:"expected_delivery"`
	ActualDelivery   *time.Time          `json:"actual_delivery,omitempty"`
	Lines            []SupplierOrderLine `json:"lines"`
}

// SupplierOrderLine references a SupplierProduct — cost comes from the entity.
type SupplierOrderLine struct {
	SupplierProductID string `json:"supplier_product_id"`
	Quantity          int    `json:"quantity"`
}

func (so SupplierOrder) IsOverdue() bool {
	return so.Status != SupplierOrderReceived && so.Status != SupplierOrderCancelled && time.Now().After(so.ExpectedDelivery)
}
