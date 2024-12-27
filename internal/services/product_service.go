package services

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/repository"
	"errors"
)

// ProductService defines the service for managing products.
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct validates and creates a new product.
func (s *ProductService) CreateProduct(product *models.Product) error {
	if err := validateProduct(product); err != nil {
		return err
	}
	return s.repo.CreateProduct(product)
}

// GetProduct retrieves a product by its ID.
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProducts retrieves all products.
func (s *ProductService) GetProducts() ([]models.Product, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct validates and updates an existing product.
func (s *ProductService) UpdateProduct(product *models.Product) error {
	if err := validateProduct(product); err != nil {
		return err
	}
	return s.repo.UpdateProduct(product)
}

// DeleteProduct removes a product by its ID.
func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

// validateProduct checks if the product fields are valid.
func validateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}
	return nil
}
