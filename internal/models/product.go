package models

import "time"

// Product represents the structure of a product in the e-commerce application.
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"` // Unique identifier for the product
	Name        string    `json:"name" gorm:"not null"` // Name of the product
	Description string    `json:"description"`          // Description of the product
	Price       float64   `json:"price" gorm:"not null"` // Price of the product
	Stock       int       `json:"stock" gorm:"not null"` // Available stock for the product
	CreatedAt   time.Time `json:"created_at"`           // Timestamp when the product was created
	UpdatedAt   time.Time `json:"updated_at"`           // Timestamp when the product was last updated
}