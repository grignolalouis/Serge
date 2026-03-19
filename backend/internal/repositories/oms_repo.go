package repositories

import (
	"serge/backend/internal/models"

	"gorm.io/gorm"
)

// OmsRepository defines methods for OMS data operations
type OmsRepository interface {
	GetOrderByNumber(orderNumber string) (*models.Order, error)
}

// gormOmsRepo implements OmsRepository using GORM
type gormOmsRepo struct {
	db *gorm.DB
}

// NewOmsRepository creates a new OMS repository
func NewOmsRepository(db *gorm.DB) OmsRepository {
	return &gormOmsRepo{db: db}
}

// GetOrderByNumber retrieves an order by its order number, including order lines
func (r *gormOmsRepo) GetOrderByNumber(orderNumber string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Lines").Where("order_number = ?", orderNumber).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
