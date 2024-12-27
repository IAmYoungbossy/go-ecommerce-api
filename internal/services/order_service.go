package services

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/repository"
	"errors"
)

// OrderService handles business logic related to orders.
type OrderService struct {
	orderRepo repository.OrderRepositoryInterface
}

// NewOrderService creates a new OrderService instance.
func NewOrderService(orderRepo repository.OrderRepositoryInterface) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

// PlaceOrder processes a new order and saves it to the database.
func (s *OrderService) PlaceOrder(order *models.Order) error {
	if err := validateOrder(order); err != nil {
		return err
	}
	return s.orderRepo.CreateOrder(order)
}

// GetOrders retrieves all orders for a specific user.
func (s *OrderService) GetOrdersByUser(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetOrdersByUser(userID)
}

// CancelOrder cancels an order if it is still in the Pending status.
func (s *OrderService) CancelOrder(orderID uint) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status != "Pending" {
		return errors.New("order cannot be canceled as it is not in Pending status")
	}
	return s.orderRepo.UpdateOrderStatus(orderID, models.OrderStatusCancelled)
}

// UpdateOrderStatus updates the status of an order (admin privilege).
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.orderRepo.UpdateOrderStatus(orderID, status)
}

// validateOrder checks if the order data is valid.
func validateOrder(order *models.Order) error {
	if order.UserID == 0 {
		return errors.New("user ID is required")
	}
	if order.ProductID == 0 {
		return errors.New("product ID is required")
	}
	if order.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	return nil
}
