package auth

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware is a middleware function that checks for a valid JWT in the request cookie.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookie
		tokenString, err := c.Cookie("auth_token")
		if err != nil {
			// If the cookie is not found, respond with unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie is missing"})
			c.Abort()
			return
		}

		// Validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token's signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorMalformed)
			}
			// Return the secret key for validation
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			// Token is invalid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract user information from the token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract the user ID from the "sub" claim (the "subject" of the token)
			userID := claims["sub"].(string)
			c.Set("userID", userID)
		} else {
			// Invalid token claims
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Token is valid, proceed to the next handler
		c.Next()
	}
}
