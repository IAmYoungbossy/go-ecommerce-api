package controllers

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController handles HTTP requests related to products.
type ProductController struct {
	ProductService *services.ProductService
}

// NewProductController creates a new ProductController instance.
func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// CreateProduct handles the creation of a new product.
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := pc.ProductService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductByID retrieves a product by its ID.
func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Convert the id to uint
	productID := uint(id)

	// Now call the service with the correct id type
	product, err := pc.ProductService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProducts retrieves all products.
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct handles the update of an existing product.
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := pc.ProductService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles the deletion of a product.
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Convert the id to uint
	productID := uint(id)

	// Now call the service with the correct id type
	if err := pc.ProductService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
