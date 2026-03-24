package domain

import "time"

type Supplier struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Region            string   `json:"region"` // e.g. "Chanthaburi", "Chiang Mai"
	Contact           string   `json:"contact"`
	LeadTimeDays      int      `json:"lead_time_days"`
	Rating            float64  `json:"rating"` // 1.0–5.0
	ProductCategories []string `json:"product_categories"`
	PaymentTerms      string   `json:"payment_terms"` // "net_15", "net_30", "cod"
}

type PurchaseOrderStatus string

const (
	POPending   PurchaseOrderStatus = "pending"
	POConfirmed PurchaseOrderStatus = "confirmed"
	POShipped   PurchaseOrderStatus = "shipped"
	POReceived  PurchaseOrderStatus = "received"
	POCancelled PurchaseOrderStatus = "cancelled"
)

type PurchaseOrder struct {
	ID               string              `json:"id"`
	SupplierID       string              `json:"supplier_id"`
	Status           PurchaseOrderStatus `json:"status"`
	OrderDate        time.Time           `json:"order_date"`
	ExpectedDelivery time.Time           `json:"expected_delivery"`
	ActualDelivery   *time.Time          `json:"actual_delivery,omitempty"`
	Lines            []PurchaseOrderLine `json:"lines"`
}

type PurchaseOrderLine struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitCost  float64 `json:"unit_cost"`
}

func (po PurchaseOrder) TotalCost() float64 {
	var total float64
	for _, l := range po.Lines {
		total += float64(l.Quantity) * l.UnitCost
	}
	return total
}

func (po PurchaseOrder) IsOverdue() bool {
	return po.Status != POReceived && po.Status != POCancelled && time.Now().After(po.ExpectedDelivery)
}
