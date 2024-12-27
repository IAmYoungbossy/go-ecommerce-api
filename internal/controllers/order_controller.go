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
func (oc *OrderController) ListOrders(c *gin.Context) {
	// Retrieve the user ID from the context (set by the middleware)
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
	// Retrieve the order ID from the request URL parameter
	orderID := c.Param("id")

	// Convert the order ID from string to uint
	oid, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		// If there is an error in parsing the order ID, return a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Convert uint64 to uint before passing it to the CancelOrder method
	oidUint := uint(oid)

	// Call the CancelOrder method in the service
	if err := oc.OrderService.CancelOrder(oidUint); err != nil {
		// If there is an error in canceling the order, return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a successful cancellation message with a 200 OK status
	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}

// UpdateOrderStatus handles the request to update the status of an order
func (oc *OrderController) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	// Log the orderID to check if it's correctly extracted from the URL
	fmt.Println("Received order ID:", orderID)

	// Define a struct to hold the status from the request body
	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}

	// Bind the request body to the statusUpdate struct
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the status
	if !isValidStatus(statusUpdate.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Valid values are: Pending, Completed, Cancelled"})
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

	// Return a success message
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
