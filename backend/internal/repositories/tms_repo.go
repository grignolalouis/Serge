package repositories

import (
	"serge/backend/internal/models"

	"gorm.io/gorm"
)

// TmsRepository defines methods for TMS data operations
type TmsRepository interface {
	GetShipmentByOrderNumber(orderNumber string) (*models.Shipment, error)
}

// gormTmsRepo implements TmsRepository using GORM
type gormTmsRepo struct {
	db *gorm.DB
}

// NewTmsRepository creates a new TMS repository
func NewTmsRepository(db *gorm.DB) TmsRepository {
	return &gormTmsRepo{db: db}
}

// GetShipmentByOrderNumber retrieves a shipment by its order number
func (r *gormTmsRepo) GetShipmentByOrderNumber(orderNumber string) (*models.Shipment, error) {
	var shipment models.Shipment
	err := r.db.Where("order_number = ?", orderNumber).First(&shipment).Error
	if err != nil {
		return nil, err
	}
	return &shipment, nil
}
