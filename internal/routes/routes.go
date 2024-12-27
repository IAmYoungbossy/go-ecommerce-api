package routes

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes and associates them with their respective controller methods.
func SetupRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
) {
	// User routes
	router.POST("/api/users/login", userController.LoginUser)
	router.POST("/api/users/logout", userController.LogoutUser)
	router.POST("/api/users/register", userController.RegisterUser)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(auth.JWTMiddleware())

	// Product routes (admin only)
	authorized.GET("/api/products", productController.GetProducts)
	authorized.POST("/api/products", productController.CreateProduct)
	authorized.PUT("/api/products/:id", productController.UpdateProduct)
	authorized.GET("/api/products/:id", productController.GetProductByID)
	authorized.DELETE("/api/products/:id", productController.DeleteProduct)

	// Order routes
	authorized.GET("/api/orders", orderController.ListOrders)
	authorized.POST("/api/orders", orderController.PlaceOrder)
	authorized.PUT("/api/orders/:id/cancel", orderController.CancelOrder)
	authorized.PUT("/api/orders/:id/status", orderController.UpdateOrderStatus)
}
