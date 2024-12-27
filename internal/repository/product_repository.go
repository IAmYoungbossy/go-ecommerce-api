package repository

import (
	"ecommerce-api/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ProductRepository defines the methods for interacting with the products in the database.
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(updatedProduct *models.Product) (*models.Product, error)
}

// productRepository implements the ProductRepository interface.
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// CreateProduct inserts a new product into the database.
func (r *productRepository) CreateProduct(product *models.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID retrieves a product by its ID.
func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates a product in the database based on the provided updated product fields.
func (r *productRepository) UpdateProduct(updatedProduct *models.Product) (*models.Product, error) {
	// Find the existing product by ID
	var existingProduct models.Product
	if err := r.db.First(&existingProduct, updatedProduct.ID).Error; err != nil {
		fmt.Println("Product not found: ", err)
		return nil, fmt.Errorf("product not found")
	}

	// Update fields only if they are provided (i.e., non-zero values)
	if updatedProduct.Name != "" {
		fmt.Println("Updating Name: ", updatedProduct.Name)
		existingProduct.Name = updatedProduct.Name
	}
	if updatedProduct.Description != "" {
		fmt.Println("Updating Description: ", updatedProduct.Description)
		existingProduct.Description = updatedProduct.Description
	}
	if updatedProduct.Price != 0 {
		fmt.Println("Updating Price: ", updatedProduct.Price)
		existingProduct.Price = updatedProduct.Price
	}
	if updatedProduct.Stock != 0 {
		fmt.Println("Updating Stock: ", updatedProduct.Stock)
		existingProduct.Stock = updatedProduct.Stock
	}

	// Save the updated product back to the database
	if err := r.db.Save(&existingProduct).Error; err != nil {
		fmt.Println("Error saving updated product: ", err)
		return nil, err
	}

	return &existingProduct, nil
}

// DeleteProduct removes a product from the database.
func (r *productRepository) DeleteProduct(id uint) error {
	result := r.db.Delete(&models.Product{}, id)
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return result.Error
}

// GetAllProducts retrieves all products from the database.
func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
