package repository

import (
	"ecommerce-api/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// UserRepository is the repository for the User model
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates and returns a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(user *models.User) error {
	// Using GORM Create method to insert the user into the database
	if err := r.DB.Create(user).Error; err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}
	return nil
}

// GetUserByEmail retrieves a user by their email.
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	// Using GORM First method to find the user by email
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID.
func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	// Using GORM First method to find the user by ID
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	// Remove the password
	user.Password = ""
	return &user, nil
}

// UpdateUser updates the user's information in the database.
func (r *UserRepository) UpdateUser(user *models.User) error {
	// Using GORM Save method to update the user's details
	if err := r.DB.Save(user).Error; err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}
	return nil
}

// DeleteUser removes a user from the database.
func (r *UserRepository) DeleteUser(id uint) error {
	// Using GORM Delete method to remove the user by their ID
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}
	return nil
}
