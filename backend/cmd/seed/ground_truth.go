package main

import (
	"fmt"
	"serge/backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// SeedGroundTruth injects deterministic data for AI agent evaluation.
// Three scenarios: Order Status, Root Cause Analysis, and Supplier Comparison.
func SeedGroundTruth(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := seedScenario1(tx); err != nil {
			return err
		}
		if err := seedScenario2(tx); err != nil {
			return err
		}
		if err := seedScenario3(tx); err != nil {
			return err
		}
		return nil
	})
}

// seedScenario1: Order Status tracking
// Order ORD-2024-0456 for MegaCorp, shipped via DHL
func seedScenario1(tx *gorm.DB) error {
	fmt.Println("\n📍 Seeding SCENARIO 1: Order Status (ORD-2024-0456)")

	// Create order
	order := models.Order{
		OrderNumber:   "ORD-2024-0456",
		CustomerName:  "MegaCorp",
		Status:        "shipped",
		PlacementDate: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		RequiredDate:  time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC),
		Lines: []models.OrderLine{
			{SKU: "SKU-WIDGET-A", Quantity: 50},
			{SKU: "SKU-GADGET-B", Quantity: 20},
		},
	}
	if err := tx.Create(&order).Error; err != nil {
		return fmt.Errorf("failed to create order ORD-2024-0456: %w", err)
	}

	// Create shipment
	shipment := models.Shipment{
		ShipmentNumber:        "SHP-2024-1892",
		OrderNumber:           "ORD-2024-0456",
		Carrier:               "DHL",
		Status:                "in transit",
		EstimatedDeliveryDate: PtrTime(time.Date(2024, 1, 17, 0, 0, 0, 0, time.UTC)),
	}
	if err := tx.Create(&shipment).Error; err != nil {
		return fmt.Errorf("failed to create shipment SHP-2024-1892: %w", err)
	}

	fmt.Println("  ✓ Order ORD-2024-0456 with 2 lines created")
	fmt.Println("  ✓ Shipment SHP-2024-1892 created")

	return nil
}

// seedScenario2: Root Cause Analysis - Supply shortage causing delay
// Order requires 50 SKU-MOTOR-X but only 12 in stock
// PO with TechParts Inc in transit
func seedScenario2(tx *gorm.DB) error {
	fmt.Println("\n📍 Seeding SCENARIO 2: Root Cause Analysis - Delay (ORD-2024-0789)")

	// Create order
	order := models.Order{
		OrderNumber:   "ORD-2024-0789",
		CustomerName:  "TechSupply Inc",
		Status:        "processing",
		PlacementDate: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		RequiredDate:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
		Lines: []models.OrderLine{
			{SKU: "SKU-MOTOR-X", Quantity: 50},
		},
	}
	if err := tx.Create(&order).Error; err != nil {
		return fmt.Errorf("failed to create order ORD-2024-0789: %w", err)
	}

	// Create WMS inventory (shortage: need 50, have 12)
	inventory := models.Inventory{
		SKU:               "SKU-MOTOR-X",
		AvailableQuantity: 12,
		ReservedQuantity:  0,
		Location:          "WAREHOUSE-A",
	}
	if err := tx.Create(&inventory).Error; err != nil {
		return fmt.Errorf("failed to create inventory for SKU-MOTOR-X: %w", err)
	}

	// Create supplier
	supplier := models.Supplier{
		Name:            "TechParts Inc",
		Category:        "electronic components",
		OnTimeRatePct:   96.0,
		AvgLeadTimeDays: 8,
		DefectRatePct:   2.5,
	}
	if err := tx.Create(&supplier).Error; err != nil {
		return fmt.Errorf("failed to create supplier TechParts Inc: %w", err)
	}

	// Create purchase order
	po := models.PurchaseOrder{
		PONumber:             "PO-2024-102",
		SupplierID:           supplier.ID,
		SKU:                  "SKU-MOTOR-X",
		Quantity:             150,
		Status:               "in transit",
		ExpectedDeliveryDate: PtrTime(time.Date(2024, 1, 18, 0, 0, 0, 0, time.UTC)),
	}
	if err := tx.Create(&po).Error; err != nil {
		return fmt.Errorf("failed to create purchase order PO-2024-102: %w", err)
	}

	fmt.Println("  ✓ Order ORD-2024-0789 with 1 line created")
	fmt.Println("  ✓ Inventory SKU-MOTOR-X created (shortage: 12/50)")
	fmt.Println("  ✓ Supplier TechParts Inc created")
	fmt.Println("  ✓ Purchase order PO-2024-102 created")

	return nil
}

// seedScenario3: Supplier comparison and performance metrics
// Three suppliers in "electronic components" category with varying performance
func seedScenario3(tx *gorm.DB) error {
	fmt.Println("\n📍 Seeding SCENARIO 3: Supplier Comparison")

	suppliers := []models.Supplier{
		{
			Name:            "Acme Electronics",
			Category:        "electronic components",
			OnTimeRatePct:   91.0,
			AvgLeadTimeDays: 12,
			DefectRatePct:   3.2,
		},
		{
			Name:            "GlobalSupply Co",
			Category:        "electronic components",
			OnTimeRatePct:   78.0,
			AvgLeadTimeDays: 15,
			DefectRatePct:   4.5,
		},
	}

	for _, s := range suppliers {
		if err := tx.Create(&s).Error; err != nil {
			return fmt.Errorf("failed to create supplier %s: %w", s.Name, err)
		}
		fmt.Printf("  ✓ %s created (%.1f%% on-time rate, %d days lead time)\n", s.Name, s.OnTimeRatePct, s.AvgLeadTimeDays)
	}

	return nil
}
