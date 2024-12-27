package models

import (
	"time"
)

// Order represents an order in the e-commerce application.
type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null;default:'Pending'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// OrderStatus represents the possible statuses of an order.
const (
	OrderStatusPending   = "Pending"
	OrderStatusCompleted = "Completed"
	OrderStatusCancelled = "Cancelled"
)
