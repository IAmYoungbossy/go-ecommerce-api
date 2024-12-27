package repository

import (
	"ecommerce-api/internal/models"

	"gorm.io/gorm"
)

// ProductRepository defines the interface for product-related database operations.
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]models.Product, error)
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

// UpdateProduct updates an existing product in the database.
func (r *productRepository) UpdateProduct(product *models.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct removes a product from the database.
func (r *productRepository) DeleteProduct(id uint) error {
	if err := r.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAllProducts retrieves all products from the database.
func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
