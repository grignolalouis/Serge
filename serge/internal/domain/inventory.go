package domain

import "time"

// ---------------------------------------------------------------------------
// Inventory domain — products, warehouses, stock
// ---------------------------------------------------------------------------

// Product is what we sell to customers. Price is defined here (not on order lines).
type Product struct {
	ID            string  `json:"id"`
	SKU           string  `json:"sku"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	UnitPrice     float64 `json:"unit_price"`
	WeightKg      float64 `json:"weight_kg"`
	ReorderPoint  int     `json:"reorder_point"`
	ShelfLifeDays int     `json:"shelf_life_days"`
	StorageType   string  `json:"storage_type"`
}

type Warehouse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type InventoryRecord struct {
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	Reserved    int       `json:"reserved"`
	LastUpdated time.Time `json:"last_updated"`
}

func (r InventoryRecord) Available() int {
	return r.Quantity - r.Reserved
}

func (r InventoryRecord) NeedsReorder(reorderPoint int) bool {
	return r.Available() < reorderPoint
}

type MovementType string

const (
	MovementInbound  MovementType = "inbound"
	MovementOutbound MovementType = "outbound"
)

type StockMovement struct {
	ID            string       `json:"id"`
	ProductID     string       `json:"product_id"`
	WarehouseID   string       `json:"warehouse_id"`
	Type          MovementType `json:"type"`
	Quantity      int          `json:"quantity"`
	ReferenceID   string       `json:"reference_id"`
	ReferenceType string       `json:"reference_type"` // "supplier_order" or "customer_order"
	CreatedAt     time.Time    `json:"created_at"`
}
