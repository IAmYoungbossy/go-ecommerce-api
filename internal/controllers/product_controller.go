package controllers

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"fmt"
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
// @Summary Create a new product
// @Description Creates a new product in the system
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 201 {object} models.Product
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Could not create product"}
// @Router /products [post]
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
// @Summary Get a product by ID
// @Description Retrieves a product by its unique ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} gin.H{"error": "Invalid product ID"}
// @Failure 500 {object} gin.H{"error": "Could not retrieve product"}
// @Router /products/{id} [get]
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
// @Summary Get all products
// @Description Retrieves all available products in the system
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} gin.H{"error": "Could not retrieve products"}
// @Router /products [get]
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct handles the update of an existing product.
// @Summary Update a product
// @Description Updates an existing product in the system
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated Product Data"
// @Success 200 {object} gin.H{"message": "Product updated successfully", "product": models.Product}
// @Failure 400 {object} gin.H{"error": "Invalid product ID"}
// @Failure 500 {object} gin.H{"error": "Could not update product"}
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID provided: ", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Set the product ID from the URL params
	product.ID = uint(id)

	// Call the service to update the product and handle returned error
	updatedProduct, err := pc.ProductService.UpdateProduct(&product)
	if err != nil {
		fmt.Println("Error updating product: ", err)
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
		}
		return
	}

	// Return the updated product with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": updatedProduct})
}

// DeleteProduct handles the deletion of a product.
// @Summary Delete a product
// @Description Deletes a product by its ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} gin.H{"message": "Product deleted successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid product ID"}
// @Failure 404 {object} gin.H{"error": "Product not found"}
// @Failure 500 {object} gin.H{"error": "Could not delete product"}
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	productID := uint(id)

	// Call the service to delete the product
	if err := pc.ProductService.DeleteProduct(productID); err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete product"})
		}
		return
	}

	// Use HTTP StatusOK for responses with a message
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
