package repositories

import (
	"serge/backend/internal/models"

	"gorm.io/gorm"
)

// SrmRepository defines methods for SRM data operations
type SrmRepository interface {
	GetSuppliersByCategory(category string) ([]models.Supplier, error)
	GetPurchaseOrdersBySKU(sku string) ([]models.PurchaseOrder, error)
}

// gormSrmRepo implements SrmRepository using GORM
type gormSrmRepo struct {
	db *gorm.DB
}

// NewSrmRepository creates a new SRM repository
func NewSrmRepository(db *gorm.DB) SrmRepository {
	return &gormSrmRepo{db: db}
}

// GetSuppliersByCategory retrieves suppliers by category
func (r *gormSrmRepo) GetSuppliersByCategory(category string) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := r.db.Where("category = ?", category).Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

// GetPurchaseOrdersBySKU retrieves purchase orders by SKU
func (r *gormSrmRepo) GetPurchaseOrdersBySKU(sku string) ([]models.PurchaseOrder, error) {
	var orders []models.PurchaseOrder
	err := r.db.Where("sku = ?", sku).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
