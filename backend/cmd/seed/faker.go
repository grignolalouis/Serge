package main

import (
	"fmt"
	"math/rand"
	"serge/backend/internal/models"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

// SeedFakeData generates and injects random data to create noise for AI agent testing.
// This helps ensure the agent can find correct information among large datasets.
func SeedFakeData(db *gorm.DB) error {
	fmt.Println("\n🔀 Generating and injecting fake data...")

	return db.Transaction(func(tx *gorm.DB) error {
		// OMS: 200 random orders
		fmt.Print("  Generating 200 OMS orders")
		if err := seedFakeOrders(tx, 200); err != nil {
			return fmt.Errorf("failed to seed fake orders: %w", err)
		}
		fmt.Println(" ✓")

		// TMS: 150 random shipments
		fmt.Print("  Generating 150 TMS shipments")
		if err := seedFakeShipments(tx, 150); err != nil {
			return fmt.Errorf("failed to seed fake shipments: %w", err)
		}
		fmt.Println(" ✓")

		// WMS: 300 random inventory items
		fmt.Print("  Generating 300 WMS inventory items")
		if err := seedFakeInventory(tx, 300); err != nil {
			return fmt.Errorf("failed to seed fake inventory: %w", err)
		}
		fmt.Println(" ✓")

		// SRM: 50 random suppliers
		fmt.Print("  Generating 50 SRM suppliers")
		supplierIDs, err := seedFakeSuppliers(tx, 50)
		if err != nil {
			return fmt.Errorf("failed to seed fake suppliers: %w", err)
		}
		fmt.Println(" ✓")

		// SRM: 100 random purchase orders
		fmt.Print("  Generating 100 SRM purchase orders")
		if err := seedFakePurchaseOrders(tx, 100, supplierIDs); err != nil {
			return fmt.Errorf("failed to seed fake purchase orders: %w", err)
		}
		fmt.Println(" ✓")

		return nil
	})
}

func seedFakeOrders(tx *gorm.DB, count int) error {
	statuses := []string{"pending", "processing", "shipped", "delivered", "cancelled"}
	skus := []string{"SKU-A", "SKU-B", "SKU-C", "SKU-D", "SKU-E"}

	for i := 0; i < count; i++ {
		order := models.Order{
			OrderNumber:   fmt.Sprintf("ORD-FAKE-%06d", i+1),
			CustomerName:  gofakeit.Company(),
			Status:        statuses[rand.Intn(len(statuses))],
			PlacementDate: gofakeit.Date().Add(time.Hour * time.Duration(rand.Intn(24))),
			RequiredDate:  gofakeit.Date().Add(time.Hour * time.Duration(rand.Intn(24))),
		}

		// Create 1-5 order lines per order
		numLines := rand.Intn(5) + 1
		for j := 0; j < numLines; j++ {
			order.Lines = append(order.Lines, models.OrderLine{
				SKU:      skus[rand.Intn(len(skus))],
				Quantity: rand.Intn(100) + 1,
			})
		}

		if err := tx.Create(&order).Error; err != nil {
			return fmt.Errorf("failed to create order %s: %w", order.OrderNumber, err)
		}
	}
	return nil
}

func seedFakeShipments(tx *gorm.DB, count int) error {
	carriers := []string{"UPS", "FedEx", "DHL", "USPS", "XPO", "J.B. Hunt"}
	statuses := []string{"pending", "in transit", "delivered", "failed", "returned"}

	// Get all existing orders to link shipments to valid orders
	var orders []models.Order
	if err := tx.Find(&orders).Error; err != nil {
		return err
	}

	// Only create shipments for available orders
	maxShipments := count
	if len(orders) < count {
		maxShipments = len(orders)
	}

	for i := 0; i < maxShipments; i++ {
		shipment := models.Shipment{
			ShipmentNumber:        fmt.Sprintf("SHP-FAKE-%06d", i+1),
			OrderNumber:           orders[i].OrderNumber,
			Carrier:               carriers[rand.Intn(len(carriers))],
			Status:                statuses[rand.Intn(len(statuses))],
			EstimatedDeliveryDate: PtrTime(gofakeit.Date()),
		}

		if err := tx.Create(&shipment).Error; err != nil {
			return fmt.Errorf("failed to create shipment %s: %w", shipment.ShipmentNumber, err)
		}
	}
	return nil
}

func seedFakeInventory(tx *gorm.DB, count int) error {
	locations := []string{"WAREHOUSE-A", "WAREHOUSE-B", "WAREHOUSE-C", "DOCK-1", "DOCK-2"}

	for i := 0; i < count; i++ {
		inventory := models.Inventory{
			SKU:               fmt.Sprintf("SKU-FAKE-%05d", i+1),
			AvailableQuantity: rand.Intn(1000) + 1,
			ReservedQuantity:  rand.Intn(100),
			Location:          locations[rand.Intn(len(locations))],
		}

		if err := tx.Create(&inventory).Error; err != nil {
			return fmt.Errorf("failed to create inventory for %s: %w", inventory.SKU, err)
		}
	}

	return nil
}

func seedFakeSuppliers(tx *gorm.DB, count int) ([]uint, error) {
	categories := []string{
		"electronic components",
		"raw materials",
		"packaging",
		"logistics",
		"manufacturing equipment",
	}

	var supplierIDs []uint

	for i := 0; i < count; i++ {
		supplier := models.Supplier{
			Name:            fmt.Sprintf("%s Supplier %d", gofakeit.Company(), i+1),
			Category:        categories[rand.Intn(len(categories))],
			OnTimeRatePct:   float64(rand.Intn(50) + 50), // 50-100%
			AvgLeadTimeDays: rand.Intn(20) + 5,            // 5-25 days
			DefectRatePct:   float64(rand.Intn(10)) / 2,   // 0-5%
		}

		if err := tx.Create(&supplier).Error; err != nil {
			return nil, fmt.Errorf("failed to create supplier %s: %w", supplier.Name, err)
		}

		supplierIDs = append(supplierIDs, supplier.ID)
	}

	return supplierIDs, nil
}

func seedFakePurchaseOrders(tx *gorm.DB, count int, supplierIDs []uint) error {
	statuses := []string{"pending", "confirmed", "in transit", "delivered", "cancelled"}
	skus := []string{"SKU-WIDGET-A", "SKU-GADGET-B", "SKU-MOTOR-X", "SKU-X", "SKU-Y", "SKU-Z"}

	if len(supplierIDs) == 0 {
		return fmt.Errorf("no supplier IDs available")
	}

	for i := 0; i < count; i++ {
		po := models.PurchaseOrder{
			PONumber:             fmt.Sprintf("PO-FAKE-%06d", i+1),
			SupplierID:           supplierIDs[rand.Intn(len(supplierIDs))],
			SKU:                  skus[rand.Intn(len(skus))],
			Quantity:             rand.Intn(500) + 10,
			Status:               statuses[rand.Intn(len(statuses))],
			ExpectedDeliveryDate: PtrTime(gofakeit.Date()),
		}

		if err := tx.Create(&po).Error; err != nil {
			return fmt.Errorf("failed to create purchase order %s: %w", po.PONumber, err)
		}
	}

	return nil
}
