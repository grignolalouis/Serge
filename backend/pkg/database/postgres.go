package database

import (
	"serge/backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config holds database configuration
type Config struct {
	DSN string
}

// Connect establishes a connection to PostgreSQL and runs migrations
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Create schemas
	schemas := []string{"oms", "tms", "wms", "srm"}
	for _, schema := range schemas {
		if err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema).Error; err != nil {
			return nil, err
		}
	}

	// AutoMigrate all schemas
	err = db.AutoMigrate(
		&models.Order{},
		&models.OrderLine{},
		&models.Shipment{},
		&models.Inventory{},
		&models.Supplier{},
		&models.PurchaseOrder{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
