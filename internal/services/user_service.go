package services

import (
	"ecommerce-api/internal/auth"
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/repository"
	"errors"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// UserService handles business logic related to users.
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser hashes the user's password and saves the user to the database.
func (s *UserService) RegisterUser(user *models.User) error {
	// Ensure email and password are provided
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// No need to hash the password here, because the BeforeSave hook will handle it
	return s.userRepo.CreateUser(user)
}

// AuthenticateUser authenticates a user and generates a JWT token
func (s *UserService) AuthenticateUser(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("Password comparison failed for user %s", user.Email)
		return "", errors.New("invalid email or password")
	}

	userIDStr := strconv.Itoa(int(user.ID))
	// Pass both userID and role to GenerateToken
	token, err := auth.GenerateToken(userIDStr, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

// LoginUser checks the user's credentials and returns an error if they are invalid.
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUser retrieves a user by their ID.
func (s *UserService) GetUserByID(email string) (*models.User, error) {
	return s.userRepo.GetUserByID(email)
}

// UpdateUser updates the user's information in the database.
func (s *UserService) UpdateUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	return s.userRepo.UpdateUser(user)
}

// DeleteUser removes a user from the database.
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
