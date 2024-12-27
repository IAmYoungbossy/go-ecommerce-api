package repository

import (
	"ecommerce-api/internal/models"

	"gorm.io/gorm"
)

// OrderRepositoryInterface defines the contract for the order repository.
type OrderRepositoryInterface interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(orderID uint) (*models.Order, error)
	GetOrdersByUser(userID uint) ([]models.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	DeleteOrder(orderID uint) error
}

// OrderRepository defines the methods for order-related database operations.
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository.
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// CreateOrder inserts a new order into the database using GORM.
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByID retrieves an order by its ID using GORM.
func (r *OrderRepository) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrdersByUser retrieves all orders for a specific user using GORM.
func (r *OrderRepository) GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrderStatus updates the status of an existing order using GORM.
func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	if err := r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder removes an order from the database using GORM.
func (r *OrderRepository) DeleteOrder(orderID uint) error {
	if err := r.db.Delete(&models.Order{}, orderID).Error; err != nil {
		return err
	}
	return nil
}
