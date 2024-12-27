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
// @Summary Register a new user
// @Description Registers a new user in the system
// @Accept  json
// @Produce  json
// @Param user body models.User true "User Information"
// @Success 201 {object} gin.H{"message": "User registered successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Could not create user"}
// @Router /users/register [post]
func (uc *UserController) RegisterUser(c *gin.Context) {
	var user models.User

	// Bind incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Register the user
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

// LoginUser handles user login
// @Summary Log in an existing user
// @Description Logs in a user and returns an authentication token
// @Accept  json
// @Produce  json
// @Param user body models.User true "Login Credentials"
// @Success 200 {object} gin.H{"token": "auth_token"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 401 {object} gin.H{"error": "Invalid credentials"}
// @Router /users/login [post]
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

// LogoutUser handles user logout
// @Summary Log out a user
// @Description Logs out a user and expires the authentication cookie
// @Success 200 {object} gin.H{"message": "Successfully logged out"}
// @Router /users/logout [post]
func (uc *UserController) LogoutUser(c *gin.Context) {
	// Expire the cookie
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

// GetUser retrieves user information
// @Summary Get user details
// @Description Retrieves details of the authenticated user
// @Produce  json
// @Success 200 {object} models.User
// @Failure 401 {object} gin.H{"error": "User ID not found in context"}
// @Failure 500 {object} gin.H{"error": "Invalid user ID type in context"}
// @Router /users/me [get]
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
// @Summary Update user details
// @Description Updates the details of the authenticated user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Updated User Information"
// @Success 200 {object} gin.H{"message": "User updated successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid input"}
// @Failure 500 {object} gin.H{"error": "Could not update user"}
// @Router /users/me [put]
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
// @Summary Delete a user
// @Description Deletes a user from the system
// @Param id path int true "User ID"
// @Success 200 {object} gin.H{"message": "User deleted successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid user ID"}
// @Failure 500 {object} gin.H{"error": "Could not delete user"}
// @Router /users/{id} [delete]
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
