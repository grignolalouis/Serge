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
	supplierProducts := must(loadScenario[[]domain.SupplierProduct]("supplier_products.json"))
	scenarioOrders := must(loadScenario[[]domain.CustomerOrder]("customer_orders.json"))
	scenarioSOs := must(loadScenario[[]domain.SupplierOrder]("supplier_orders.json"))
	scenarioShipments := must(loadScenario[[]domain.Shipment]("shipments.json"))
	scenarioEvents := must(loadScenario[[]domain.TrackingEvent]("tracking_events.json"))
	scenarioMovements := must(loadScenario[[]domain.StockMovement]("stock_movements.json"))
	inventory := must(loadScenario[[]domain.InventoryRecord]("inventory.json"))

	start := time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC)

	genOrders := generateCustomerOrders(rng, customers, products, 470, start, end)
	genSOs := generateSupplierOrders(rng, suppliers, supplierProducts, 42, start, end)
	genShipments, genEvents := generateShipments(rng, genOrders, carriers, products)
	genOrderMvts := generateCustomerOrderMovements(genOrders)
	genSOMvts := generateSupplierOrderMovements(genSOs, supplierProducts)

	allOrders := append(scenarioOrders, genOrders...)
	allSOs := append(scenarioSOs, genSOs...)
	allShipments := append(scenarioShipments, genShipments...)
	allEvents := append(scenarioEvents, genEvents...)
	allMvts := append(scenarioMovements, genOrderMvts...)
	allMvts = append(allMvts, genSOMvts...)

	writeJSON(filepath.Join(outDir, "suppliers.json"), suppliers)
	writeJSON(filepath.Join(outDir, "products.json"), products)
	writeJSON(filepath.Join(outDir, "warehouses.json"), warehouses)
	writeJSON(filepath.Join(outDir, "customers.json"), customers)
	writeJSON(filepath.Join(outDir, "carriers.json"), carriers)
	writeJSON(filepath.Join(outDir, "supplier_products.json"), supplierProducts)
	writeJSON(filepath.Join(outDir, "inventory.json"), inventory)
	writeJSON(filepath.Join(outDir, "customer_orders.json"), allOrders)
	writeJSON(filepath.Join(outDir, "supplier_orders.json"), allSOs)
	writeJSON(filepath.Join(outDir, "shipments.json"), allShipments)
	writeJSON(filepath.Join(outDir, "tracking_events.json"), allEvents)
	writeJSON(filepath.Join(outDir, "stock_movements.json"), allMvts)

	fmt.Printf("Seed data written to %s/\n", outDir)
	fmt.Printf("  Customer Orders:  %d (scenario: %d + generated: %d)\n", len(allOrders), len(scenarioOrders), len(genOrders))
	fmt.Printf("  Supplier Orders:  %d (scenario: %d + generated: %d)\n", len(allSOs), len(scenarioSOs), len(genSOs))
	fmt.Printf("  Shipments:        %d (scenario: %d + generated: %d)\n", len(allShipments), len(scenarioShipments), len(genShipments))
	fmt.Printf("  Tracking Events:  %d (scenario: %d + generated: %d)\n", len(allEvents), len(scenarioEvents), len(genEvents))
	fmt.Printf("  Stock Movements:  %d (scenario: %d + generated: %d)\n", len(allMvts), len(scenarioMovements), len(genOrderMvts)+len(genSOMvts))
}

func generateCustomerOrders(rng *rand.Rand, customers []domain.Customer, products []domain.Product, count int, start, end time.Time) []domain.CustomerOrder {
	priorities := []domain.Priority{domain.PriorityNormal, domain.PriorityNormal, domain.PriorityNormal, domain.PriorityHigh, domain.PriorityUrgent}
	notes := []string{"", "", "", "", "", "", "", "", "", "Deliver before 9am", "Call before delivery", "Leave at loading dock B", "Weekly standing order", "Fragile — handle with care"}

	orders := make([]domain.CustomerOrder, 0, count)
	for i := 0; i < count; i++ {
		cust := customers[rng.Intn(len(customers))]
		orderDate := randomDate(rng, start, end)
		requiredDate := orderDate.Add(time.Duration(3+rng.Intn(5)) * 24 * time.Hour)
		shippedDate := orderDate.Add(time.Duration(1+rng.Intn(2)) * 24 * time.Hour)
		deliveredDate := shippedDate.Add(time.Duration(1+rng.Intn(2)) * 24 * time.Hour)

		numLines := 1 + rng.Intn(3)
		var lines []domain.CustomerOrderLine
		used := map[string]bool{}
		for j := 0; j < numLines; j++ {
			p := products[rng.Intn(len(products))]
			if used[p.ID] {
				continue
			}
			used[p.ID] = true
			lines = append(lines, domain.CustomerOrderLine{
				ProductID: p.ID,
				Quantity:  quantityForSegment(rng, cust.Segment),
			})
		}
		if len(lines) == 0 {
			continue
		}

		orders = append(orders, domain.CustomerOrder{
			ID:            fmt.Sprintf("ORD-%03d", 100+i),
			CustomerID:    cust.ID,
			Status:        domain.CustomerOrderDelivered,
			OrderDate:     orderDate,
			RequiredDate:  requiredDate,
			ShippedDate:   &shippedDate,
			DeliveredDate: &deliveredDate,
			Priority:      priorities[rng.Intn(len(priorities))],
			Notes:         notes[rng.Intn(len(notes))],
			Lines:         lines,
		})
	}
	return orders
}

func generateSupplierOrders(rng *rand.Rand, suppliers []domain.Supplier, sps []domain.SupplierProduct, count int, start, end time.Time) []domain.SupplierOrder {
	supplierSPs := map[string][]domain.SupplierProduct{}
	for _, sp := range sps {
		supplierSPs[sp.SupplierID] = append(supplierSPs[sp.SupplierID], sp)
	}

	sos := make([]domain.SupplierOrder, 0, count)
	for i := 0; i < count; i++ {
		sup := suppliers[rng.Intn(len(suppliers))]
		prods := supplierSPs[sup.ID]
		if len(prods) == 0 {
			continue
		}

		orderDate := randomDate(rng, start, end)
		expected := orderDate.Add(time.Duration(sup.LeadTimeDays) * 24 * time.Hour)
		delay := rng.Intn(3)
		if sup.Rating >= 4.0 {
			delay = 0
		}
		actual := expected.Add(time.Duration(delay) * 24 * time.Hour)

		numLines := 1 + rng.Intn(2)
		var lines []domain.SupplierOrderLine
		used := map[string]bool{}
		for j := 0; j < numLines; j++ {
			sp := prods[rng.Intn(len(prods))]
			if used[sp.ID] {
				continue
			}
			used[sp.ID] = true
			lines = append(lines, domain.SupplierOrderLine{
				SupplierProductID: sp.ID,
				Quantity:          20 + rng.Intn(130),
			})
		}
		if len(lines) == 0 {
			continue
		}

		sos = append(sos, domain.SupplierOrder{
			ID:               fmt.Sprintf("PO-%03d", 100+i),
			SupplierID:       sup.ID,
			Status:           domain.SupplierOrderReceived,
			OrderDate:        orderDate,
			ExpectedDelivery: expected,
			ActualDelivery:   &actual,
			Lines:            lines,
		})
	}
	return sos
}

func generateShipments(rng *rand.Rand, orders []domain.CustomerOrder, carriers []domain.Carrier, products []domain.Product) ([]domain.Shipment, []domain.TrackingEvent) {
	productMap := map[string]domain.Product{}
	for _, p := range products {
		productMap[p.ID] = p
	}

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

		needsRefrigeration := false
		var totalWeight float64
		for _, line := range ord.Lines {
			p := productMap[line.ProductID]
			totalWeight += p.WeightKg * float64(line.Quantity)
			if p.StorageType == "refrigerated" {
				needsRefrigeration = true
			}
		}

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
		shipID := fmt.Sprintf("SHP-%s", ord.ID[4:])

		shipments = append(shipments, domain.Shipment{
			ID:               shipID,
			CustomerOrderID:  ord.ID,
			CarrierID:        carrier.ID,
			Status:           domain.ShipmentDelivered,
			TrackingNumber:   fmt.Sprintf("%s-%s", carrier.Name[:3], ord.ID[4:]),
			Origin:           "Bangkok Central Warehouse",
			Destination:      "Bangkok",
			WeightKg:         totalWeight,
			ShippedAt:        &shippedAt,
			EstimatedArrival: &eta,
			DeliveredAt:      ord.DeliveredDate,
		})

		events = append(events, generateTrackingEvents(rng, shipID, shippedAt, ord.DeliveredDate)...)
	}
	return shipments, events
}

func generateTrackingEvents(rng *rand.Rand, shipmentID string, shipped time.Time, delivered *time.Time) []domain.TrackingEvent {
	events := []domain.TrackingEvent{
		{ShipmentID: shipmentID, Timestamp: shipped.Add(30 * time.Minute), Location: "Bang Na, Bangkok", Status: "picked_up", Description: "Package picked up from Bangkok Central Warehouse"},
		{ShipmentID: shipmentID, Timestamp: shipped.Add(time.Duration(4+rng.Intn(4)) * time.Hour), Location: "Bangkok", Status: "in_transit", Description: "Departed Bangkok sorting facility"},
	}
	if delivered != nil {
		events = append(events, domain.TrackingEvent{ShipmentID: shipmentID, Timestamp: *delivered, Location: "Bangkok", Status: "delivered", Description: "Delivered — signed by recipient"})
	}
	return events
}

func generateCustomerOrderMovements(orders []domain.CustomerOrder) []domain.StockMovement {
	var mvts []domain.StockMovement
	idx := 100
	for _, ord := range orders {
		if ord.ShippedDate == nil {
			continue
		}
		for _, line := range ord.Lines {
			mvts = append(mvts, domain.StockMovement{
				ID: fmt.Sprintf("SM-%04d", idx), ProductID: line.ProductID, WarehouseID: "WH-BKK",
				Type: domain.MovementOutbound, Quantity: line.Quantity,
				ReferenceID: ord.ID, ReferenceType: "customer_order", CreatedAt: *ord.ShippedDate,
			})
			idx++
		}
	}
	return mvts
}

func generateSupplierOrderMovements(sos []domain.SupplierOrder, sps []domain.SupplierProduct) []domain.StockMovement {
	spMap := map[string]domain.SupplierProduct{}
	for _, sp := range sps {
		spMap[sp.ID] = sp
	}

	var mvts []domain.StockMovement
	idx := 2000
	for _, so := range sos {
		if so.ActualDelivery == nil {
			continue
		}
		for _, line := range so.Lines {
			sp := spMap[line.SupplierProductID]
			mvts = append(mvts, domain.StockMovement{
				ID: fmt.Sprintf("SM-%04d", idx), ProductID: sp.ProductID, WarehouseID: "WH-BKK",
				Type: domain.MovementInbound, Quantity: line.Quantity,
				ReferenceID: so.ID, ReferenceType: "supplier_order", CreatedAt: *so.ActualDelivery,
			})
			idx++
		}
	}
	return mvts
}

func quantityForSegment(rng *rand.Rand, segment domain.CustomerSegment) int {
	if segment == domain.SegmentWholesale {
		return 5 + rng.Intn(25)
	}
	return 2 + rng.Intn(15)
}

func randomDate(rng *rand.Rand, start, end time.Time) time.Time {
	diff := end.Sub(start)
	d := start.Add(time.Duration(rng.Int63n(int64(diff))))
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
