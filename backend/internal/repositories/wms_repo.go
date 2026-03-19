package repositories

import (
	"serge/backend/internal/models"

	"gorm.io/gorm"
)

// WmsRepository defines methods for WMS data operations
type WmsRepository interface {
	GetInventoryBySKU(sku string) (*models.Inventory, error)
}

// gormWmsRepo implements WmsRepository using GORM
type gormWmsRepo struct {
	db *gorm.DB
}

// NewWmsRepository creates a new WMS repository
func NewWmsRepository(db *gorm.DB) WmsRepository {
	return &gormWmsRepo{db: db}
}

// GetInventoryBySKU retrieves inventory by SKU
func (r *gormWmsRepo) GetInventoryBySKU(sku string) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Where("sku = ?", sku).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}
