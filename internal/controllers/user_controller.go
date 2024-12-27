package controllers

import (
	"ecommerce-api/internal/models"
	"ecommerce-api/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related operations.
type UserController struct {
	UserService *services.UserService
}

// NewUserController creates a new UserController instance.
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := uc.UserService.RegisterUser(&user); err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Could not create user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Set Cookie with descriptive names and correct argument order
func (uc *UserController) LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := uc.UserService.AuthenticateUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	domain := ""
	cookiePath := "/"
	secureFlag := false
	httpOnlyFlag := true
	cookieValue := token
	maxAgeInSeconds := 3600
	cookieName := "access_token"

	c.SetCookie(cookieName, cookieValue, maxAgeInSeconds, cookiePath, domain, secureFlag, httpOnlyFlag)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LogoutUser handles the logout functionality.
func (uc *UserController) LogoutUser(c *gin.Context) {
	// Expire the cookie
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

// GetUser retrieves user information
func (uc *UserController) GetUser(c *gin.Context) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to string
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in context"})
		return
	}

	user, err := uc.UserService.GetUserByID(userIDStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user information
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := uc.UserService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	// Get userID from URL parameter
	userIDStr := c.Param("id")

	// Convert userID from string to uint
	userID, err := strconv.ParseUint(userIDStr, 10, 32) // 32-bit unsigned integer (uint)
	if err != nil {
		// If conversion fails, return an error response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the DeleteUser service function with the converted uint
	if err := uc.UserService.DeleteUser(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
