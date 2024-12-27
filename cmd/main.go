package main

import (
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/controllers"
	"ecommerce-api/internal/database"
	"ecommerce-api/internal/logger"
	"ecommerce-api/internal/models" // Import the models package
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/routes"
	"ecommerce-api/internal/services"

	"github.com/gin-gonic/gin"
)

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
		&models.User{},    // Migration for User model
		&models.Order{},   // Migration for Order model (you need to create Order model)
		&models.Product{}, // Migration for Product model (you need to create Product model)
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
