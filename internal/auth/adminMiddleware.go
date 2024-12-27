package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminMiddleware checks if the user has admin privileges.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve claims from the context (set by JWTMiddleware)
		userRole, exists := c.Get("userRole")
		if !exists || userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: Admins only"})
			c.Abort()
			return
		}

		// Continue to the next middleware/handler
		c.Next()
	}
}
