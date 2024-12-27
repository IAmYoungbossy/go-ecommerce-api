package controllers

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"fmt"
	"log"
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
// @Summary Place a new order
// @Description Create a new order with the provided details for the authenticated user
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order object containing items and shipping details"
// @Success 201 {object} models.Order "Successfully created order"
// @Failure 400 {object} gin.H "Invalid input or malformed request body"
// @Failure 401 {object} gin.H "User not authenticated or invalid authentication token"
// @Failure 500 {object} gin.H "Internal server error while processing the order"
// @Security ApiKeyAuth
// @Router /orders [post]
func (oc *OrderController) PlaceOrder(c *gin.Context) {
	var order models.Order

	// Retrieve the user ID from the context (set by the authentication middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to uint (assuming userID is a string in the token)
	uid, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert int to uint before assigning to order.UserID
	uidUint := uint(uid)
	order.UserID = uidUint

	// Set the status to "Pending"
	order.Status = models.OrderStatusPending

	// Bind the request body to the order struct
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Log the order to ensure UserID is not overwritten
	log.Printf("Order details: %+v", order)

	// Call the service to place the order
	if err := oc.OrderService.PlaceOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created order
	c.JSON(http.StatusCreated, order)
}

// ListOrders handles the request to list all orders for a specific user.
// @Summary List all orders for a user
// @Description Retrieve all orders placed by the authenticated user
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 401 {object} gin.H{"error": "User not authenticated"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders [get]
func (oc *OrderController) ListOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	uid, err := strconv.Atoi(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	uidUint := uint(uid)
	orders, err := oc.OrderService.GetOrdersByUser(uidUint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// CancelOrder handles the request to cancel an order
// @Summary Cancel an order
// @Description Cancel a specific order by ID
// @Tags Orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} gin.H{"message": "Order canceled successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid order ID"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders/{id} [delete]
func (oc *OrderController) CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	oid, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	oidUint := uint(oid)
	if err := oc.OrderService.CancelOrder(oidUint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// UpdateOrderStatus handles the request to update the status of an order
// @Summary Update order status
// @Description Update the status of a specific order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body string true "Order status (Pending, Completed, Cancelled)"
// @Success 200 {object} gin.H{"message": "Order status updated"}
// @Failure 400 {object} gin.H{"error": "Invalid input or status"}
// @Failure 500 {object} gin.H{"error": "Internal server error"}
// @Router /orders/{id}/status [put]
func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	fmt.Println("Received order ID:", orderID)

	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if !isValidStatus(statusUpdate.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Valid values are: Pending, Completed, Cancelled"})
		return
	}

	oid, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	oidUint := uint(oid)
	if err := oc.OrderService.UpdateOrderStatus(oidUint, statusUpdate.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}

// Helper function to validate status
func isValidStatus(status string) bool {
	validStatuses := []string{"Pending", "Completed", "Cancelled"}
	for _, s := range validStatuses {
		if status == s {
			return true
		}
	}
	return false
}
