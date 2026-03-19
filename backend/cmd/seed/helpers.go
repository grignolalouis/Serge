package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// PtrTime creates a pointer to a time.Time value
func PtrTime(t time.Time) *time.Time {
	return &t
}

// ClearDatabase removes all data from all tables (use with caution - development only).
// Respects foreign key dependencies by deleting in correct order.
func ClearDatabase(db *gorm.DB) error {
	fmt.Println("⚠️  Clearing all database tables...")

	return db.Transaction(func(tx *gorm.DB) error {
		// Delete in reverse order of foreign key dependencies
		tables := []string{
			"srm.purchase_orders",
			"srm.suppliers",
			"wms.inventory",
			"tms.shipments",
			"oms.order_lines",
			"oms.orders",
		}

		for _, table := range tables {
			if err := tx.Exec("DELETE FROM " + table).Error; err != nil {
				return fmt.Errorf("failed to clear table %s: %w", table, err)
			}
		}

		fmt.Println("✓ Database cleared")
		return nil
	})
}
