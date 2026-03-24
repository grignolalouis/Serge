package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/serge-music/serge/internal/domain"
)

//go:embed scenarios/*.json
var scenarioFS embed.FS

func main() {
	rng := rand.New(rand.NewSource(42))

	outDir := "data"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	suppliers := must(loadScenario[[]domain.Supplier]("suppliers.json"))
	products := must(loadScenario[[]domain.Product]("products.json"))
	warehouses := must(loadScenario[[]domain.Warehouse]("warehouses.json"))
	customers := must(loadScenario[[]domain.Customer]("customers.json"))
	carriers := must(loadScenario[[]domain.Carrier]("carriers.json"))
	scenarioOrders := must(loadScenario[[]domain.Order]("orders.json"))
	scenarioPOs := must(loadScenario[[]domain.PurchaseOrder]("purchase_orders.json"))
	scenarioShipments := must(loadScenario[[]domain.Shipment]("shipments.json"))
	scenarioEvents := must(loadScenario[[]domain.TrackingEvent]("tracking_events.json"))
	scenarioMovements := must(loadScenario[[]domain.StockMovement]("stock_movements.json"))
	inventory := must(loadScenario[[]domain.InventoryRecord]("inventory.json"))

	// ── Generate additional data (IDs >= 100) ───────────────────────────
	startDate := time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC)

	genOrders := generateOrders(rng, customers, products, 470, startDate, endDate)
	genPOs := generatePOs(rng, suppliers, products, 42, startDate, endDate)
	genShipments, genEvents := generateShipments(rng, genOrders, carriers, products)
	genOrderMovements := generateOrderMovements(rng, genOrders)
	genPOMovements := generatePOMovements(rng, genPOs)

	allOrders := append(scenarioOrders, genOrders...)
	allPOs := append(scenarioPOs, genPOs...)
	allShipments := append(scenarioShipments, genShipments...)
	allEvents := append(scenarioEvents, genEvents...)
	allMovements := append(scenarioMovements, genOrderMovements...)
	allMovements = append(allMovements, genPOMovements...)

	// ── Write output ────────────────────────────────────────────────────
	writeJSON(filepath.Join(outDir, "suppliers.json"), suppliers)
	writeJSON(filepath.Join(outDir, "products.json"), products)
	writeJSON(filepath.Join(outDir, "warehouses.json"), warehouses)
	writeJSON(filepath.Join(outDir, "customers.json"), customers)
	writeJSON(filepath.Join(outDir, "carriers.json"), carriers)
	writeJSON(filepath.Join(outDir, "inventory.json"), inventory)
	writeJSON(filepath.Join(outDir, "orders.json"), allOrders)
	writeJSON(filepath.Join(outDir, "purchase_orders.json"), allPOs)
	writeJSON(filepath.Join(outDir, "shipments.json"), allShipments)
	writeJSON(filepath.Join(outDir, "tracking_events.json"), allEvents)
	writeJSON(filepath.Join(outDir, "stock_movements.json"), allMovements)

	fmt.Printf("Seed data written to %s/\n", outDir)
	fmt.Printf("  Orders:          %d (scenario: %d + generated: %d)\n", len(allOrders), len(scenarioOrders), len(genOrders))
	fmt.Printf("  Purchase Orders: %d (scenario: %d + generated: %d)\n", len(allPOs), len(scenarioPOs), len(genPOs))
	fmt.Printf("  Shipments:       %d (scenario: %d + generated: %d)\n", len(allShipments), len(scenarioShipments), len(genShipments))
	fmt.Printf("  Tracking Events: %d (scenario: %d + generated: %d)\n", len(allEvents), len(scenarioEvents), len(genEvents))
	fmt.Printf("  Stock Movements: %d (scenario: %d + generated: %d)\n", len(allMovements), len(scenarioMovements), len(genOrderMovements)+len(genPOMovements))
}

func generateOrders(rng *rand.Rand, customers []domain.Customer, products []domain.Product, count int, start, end time.Time) []domain.Order {
	priorities := []domain.Priority{domain.PriorityNormal, domain.PriorityNormal, domain.PriorityNormal, domain.PriorityHigh, domain.PriorityUrgent}
	notes := []string{
		"", "", "", "", "", "", "", "", "", // 90% no notes
		"Deliver before 9am",
		"Call before delivery",
		"Leave at loading dock B",
		"Weekly standing order",
		"Fragile — handle with care",
	}

	orders := make([]domain.Order, 0, count)
	for i := 0; i < count; i++ {
		cust := customers[rng.Intn(len(customers))]
		orderDate := randomDate(rng, start, end)
		requiredDate := orderDate.Add(time.Duration(3+rng.Intn(5)) * 24 * time.Hour)
		shippedDate := orderDate.Add(time.Duration(1+rng.Intn(2)) * 24 * time.Hour)
		deliveredDate := shippedDate.Add(time.Duration(1+rng.Intn(2)) * 24 * time.Hour)

		// 1–3 line items, no duplicate products
		numLines := 1 + rng.Intn(3)
		var lines []domain.OrderLine
		used := map[string]bool{}
		for j := 0; j < numLines; j++ {
			p := products[rng.Intn(len(products))]
			if used[p.ID] {
				continue
			}
			used[p.ID] = true
			qty := quantityForSegment(rng, cust.Segment)
			lines = append(lines, domain.OrderLine{
				ProductID: p.ID,
				Quantity:  qty,
				UnitPrice: p.UnitPrice,
			})
		}
		if len(lines) == 0 {
			continue
		}

		note := notes[rng.Intn(len(notes))]

		orders = append(orders, domain.Order{
			ID:            fmt.Sprintf("ORD-%03d", 100+i),
			CustomerID:    cust.ID,
			Status:        domain.OrderDelivered,
			OrderDate:     orderDate,
			RequiredDate:  requiredDate,
			ShippedDate:   &shippedDate,
			DeliveredDate: &deliveredDate,
			Priority:      priorities[rng.Intn(len(priorities))],
			Notes:         note,
			Lines:         lines,
		})
	}
	return orders
}

func generatePOs(rng *rand.Rand, suppliers []domain.Supplier, products []domain.Product, count int, start, end time.Time) []domain.PurchaseOrder {
	// Build supplier → products map
	supplierProducts := map[string][]domain.Product{}
	for _, p := range products {
		for _, sid := range p.SupplierIDs {
			supplierProducts[sid] = append(supplierProducts[sid], p)
		}
	}

	pos := make([]domain.PurchaseOrder, 0, count)
	for i := 0; i < count; i++ {
		sup := suppliers[rng.Intn(len(suppliers))]
		prods := supplierProducts[sup.ID]
		if len(prods) == 0 {
			continue
		}

		orderDate := randomDate(rng, start, end)
		expected := orderDate.Add(time.Duration(sup.LeadTimeDays) * 24 * time.Hour)

		// Supplier reliability affects actual delivery
		delay := rng.Intn(3) // 0-2 days delay
		if sup.Rating >= 4.0 {
			delay = 0 // Reliable suppliers deliver on time
		}
		actualDelivery := expected.Add(time.Duration(delay) * 24 * time.Hour)

		// 1-2 line items
		numLines := 1 + rng.Intn(2)
		var lines []domain.PurchaseOrderLine
		used := map[string]bool{}
		for j := 0; j < numLines; j++ {
			p := prods[rng.Intn(len(prods))]
			if used[p.ID] {
				continue
			}
			used[p.ID] = true
			qty := 20 + rng.Intn(130) // 20-150 units
			lines = append(lines, domain.PurchaseOrderLine{
				ProductID: p.ID,
				Quantity:  qty,
				UnitCost:  p.UnitPrice * 0.6, // ~60% of retail
			})
		}
		if len(lines) == 0 {
			continue
		}

		pos = append(pos, domain.PurchaseOrder{
			ID:               fmt.Sprintf("PO-%03d", 100+i),
			SupplierID:       sup.ID,
			Status:           domain.POReceived,
			OrderDate:        orderDate,
			ExpectedDelivery: expected,
			ActualDelivery:   &actualDelivery,
			Lines:            lines,
		})
	}
	return pos
}

func generateShipments(rng *rand.Rand, orders []domain.Order, carriers []domain.Carrier, products []domain.Product) ([]domain.Shipment, []domain.TrackingEvent) {
	productMap := map[string]domain.Product{}
	for _, p := range products {
		productMap[p.ID] = p
	}

	// Separate carriers by type
	var refCarriers, groundCarriers []domain.Carrier
	for _, c := range carriers {
		if c.Type == "refrigerated" {
			refCarriers = append(refCarriers, c)
		} else {
			groundCarriers = append(groundCarriers, c)
		}
	}

	var shipments []domain.Shipment
	var events []domain.TrackingEvent

	for _, ord := range orders {
		if ord.ShippedDate == nil {
			continue
		}

		// Determine if any product needs refrigeration
		needsRefrigeration := false
		var totalWeight float64
		for _, line := range ord.Lines {
			p := productMap[line.ProductID]
			totalWeight += p.WeightKg * float64(line.Quantity)
			if p.StorageType == "refrigerated" {
				needsRefrigeration = true
			}
		}

		// Pick carrier
		var carrier domain.Carrier
		if needsRefrigeration && len(refCarriers) > 0 {
			carrier = refCarriers[rng.Intn(len(refCarriers))]
		} else if len(groundCarriers) > 0 {
			carrier = groundCarriers[rng.Intn(len(groundCarriers))]
		} else {
			carrier = carriers[rng.Intn(len(carriers))]
		}

		shippedAt := *ord.ShippedDate
		eta := shippedAt.Add(time.Duration(carrier.AvgTransitDays) * 24 * time.Hour)
		deliveredAt := ord.DeliveredDate

		shipID := fmt.Sprintf("SHP-%s", ord.ID[4:]) // ORD-100 → SHP-100

		shipments = append(shipments, domain.Shipment{
			ID:               shipID,
			OrderID:          ord.ID,
			CarrierID:        carrier.ID,
			Status:           domain.ShipmentDelivered,
			TrackingNumber:   fmt.Sprintf("%s-%s", carrier.Name[:3], ord.ID[4:]),
			Origin:           "Bangkok Central Warehouse",
			Destination:      "Bangkok",
			WeightKg:         totalWeight,
			ShippedAt:        &shippedAt,
			EstimatedArrival: &eta,
			DeliveredAt:      deliveredAt,
		})

		// Generate 3-4 tracking events
		events = append(events, generateTrackingEvents(rng, shipID, shippedAt, deliveredAt)...)
	}
	return shipments, events
}

func generateTrackingEvents(rng *rand.Rand, shipmentID string, shipped time.Time, delivered *time.Time) []domain.TrackingEvent {
	events := []domain.TrackingEvent{
		{
			ShipmentID:  shipmentID,
			Timestamp:   shipped.Add(30 * time.Minute),
			Location:    "Bang Na, Bangkok",
			Status:      "picked_up",
			Description: "Package picked up from Bangkok Central Warehouse",
		},
		{
			ShipmentID:  shipmentID,
			Timestamp:   shipped.Add(time.Duration(4+rng.Intn(4)) * time.Hour),
			Location:    "Bangkok",
			Status:      "in_transit",
			Description: "Departed Bangkok sorting facility",
		},
	}

	if delivered != nil {
		events = append(events, domain.TrackingEvent{
			ShipmentID:  shipmentID,
			Timestamp:   *delivered,
			Location:    "Bangkok",
			Status:      "delivered",
			Description: "Delivered — signed by recipient",
		})
	}

	return events
}

func generateOrderMovements(rng *rand.Rand, orders []domain.Order) []domain.StockMovement {
	var movements []domain.StockMovement
	idx := 100
	for _, ord := range orders {
		if ord.ShippedDate == nil {
			continue
		}
		for _, line := range ord.Lines {
			movements = append(movements, domain.StockMovement{
				ID:            fmt.Sprintf("SM-%04d", idx),
				ProductID:     line.ProductID,
				WarehouseID:   "WH-BKK",
				Type:          domain.MovementOutbound,
				Quantity:      line.Quantity,
				ReferenceID:   ord.ID,
				ReferenceType: "sales_order",
				CreatedAt:     *ord.ShippedDate,
			})
			idx++
		}
	}
	return movements
}

func generatePOMovements(rng *rand.Rand, pos []domain.PurchaseOrder) []domain.StockMovement {
	var movements []domain.StockMovement
	idx := 2000
	for _, po := range pos {
		if po.ActualDelivery == nil {
			continue
		}
		for _, line := range po.Lines {
			movements = append(movements, domain.StockMovement{
				ID:            fmt.Sprintf("SM-%04d", idx),
				ProductID:     line.ProductID,
				WarehouseID:   "WH-BKK",
				Type:          domain.MovementInbound,
				Quantity:      line.Quantity,
				ReferenceID:   po.ID,
				ReferenceType: "purchase_order",
				CreatedAt:     *po.ActualDelivery,
			})
			idx++
		}
	}
	return movements
}

func quantityForSegment(rng *rand.Rand, segment domain.CustomerSegment) int {
	if segment == domain.SegmentWholesale {
		return 5 + rng.Intn(25) // 5-30
	}
	return 2 + rng.Intn(15) // 2-17
}

func randomDate(rng *rand.Rand, start, end time.Time) time.Time {
	diff := end.Sub(start)
	offset := time.Duration(rng.Int63n(int64(diff)))
	d := start.Add(offset)
	return time.Date(d.Year(), d.Month(), d.Day(), 8, 0, 0, 0, time.UTC)
}

func loadScenario[T any](name string) (T, error) {
	var v T
	data, err := scenarioFS.ReadFile("scenarios/" + name)
	if err != nil {
		return v, fmt.Errorf("read %s: %w", name, err)
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return v, fmt.Errorf("parse %s: %w", name, err)
	}
	return v, nil
}

func writeJSON(path string, v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal %s: %v\n", path, err)
		os.Exit(1)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write %s: %v\n", path, err)
		os.Exit(1)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
	return v
}
