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
	authorizedAdmin := authorized.Group("/")
	authorizedAdmin.Use(auth.AdminMiddleware())
	authorizedAdmin.GET("/api/products", productController.GetProducts)
	authorizedAdmin.POST("/api/products", productController.CreateProduct)
	authorizedAdmin.PUT("/api/products/:id", productController.UpdateProduct)
	authorizedAdmin.GET("/api/products/:id", productController.GetProductByID)
	authorizedAdmin.DELETE("/api/products/:id", productController.DeleteProduct)

	// Order routes
	authorized.GET("/api/users", userController.GetUser)
	authorized.GET("/api/orders", orderController.ListOrders)
	authorized.POST("/api/orders", orderController.PlaceOrder)
	authorized.PUT("/api/orders/:id/cancel", orderController.CancelOrder)
	authorized.PUT("/api/orders/:id/status", orderController.UpdateOrderStatus)
}
