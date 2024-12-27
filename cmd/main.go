package main

import (
	_ "ecommerce-api/docs"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/controllers"
	"ecommerce-api/internal/database"
	"ecommerce-api/internal/logger"
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/routes"
	"ecommerce-api/internal/services"

	"github.com/gin-gonic/gin"
)

// @title E-commerce API
// @version 1.0
// @description This is a sample e-commerce API for managing users, products, and orders.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:9543
// @BasePath /api
func main() {
	// Initialize the logger
	logger.InitLogger()

	// Load configuration settings
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Error loading configuration: " + err.Error())
	}

	// Initialize database connection
	database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db := database.GetDB()

	// Run migrations for all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
	)
	if err != nil {
		logger.Fatal("Error running migrations: " + err.Error())
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	orderService := services.NewOrderService(orderRepo)
	productService := services.NewProductService(productRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	orderController := controllers.NewOrderController(orderService)
	productController := controllers.NewProductController(productService)

	// Initialize Gin router
	router := gin.Default()

	// Set up routes with the controllers
	routes.SetupRoutes(router, userController, productController, orderController)

	// Start the server
	if err := router.Run(cfg.ServerAddress); err != nil {
		logger.Fatal("Error starting server: " + err.Error())
	}
}
