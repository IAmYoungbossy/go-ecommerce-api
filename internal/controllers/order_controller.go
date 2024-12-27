package controllers

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// OrderController handles HTTP requests related to orders
type OrderController struct {
	OrderService *services.OrderService
}

// NewOrderController creates a new instance of OrderController
func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// PlaceOrder handles the request to place a new order
func (oc *OrderController) PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := oc.OrderService.PlaceOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// ListOrders handles the request to list all orders for a specific user
func (oc *OrderController) ListOrders(c *gin.Context) {
	userID := c.Param("userId")
	uid, err := strconv.Atoi(userID) // Convert userId to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert int to uint before passing to GetOrdersByUser
	uidUint := uint(uid)

	// Pass the uint to the service method
	orders, err := oc.OrderService.GetOrdersByUser(uidUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// CancelOrder handles the request to cancel an order
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID := c.Param("orderId")
	oid, err := strconv.ParseUint(orderID, 10, 32) // 10 is the base (decimal), 32 is the bit size for uint32
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Convert uint64 to uint before passing it to CancelOrder method
	oidUint := uint(oid)

	// Pass the uint value to the CancelOrder method
	if err := oc.OrderService.CancelOrder(oidUint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UpdateOrderStatus handles the request to update the status of an order
func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("orderId")
	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert orderID from string to int
	oid, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Convert from int to uint before passing to the service
	oidUint := uint(oid)

	// Update the order status
	if err := oc.OrderService.UpdateOrderStatus(oidUint, statusUpdate.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}
