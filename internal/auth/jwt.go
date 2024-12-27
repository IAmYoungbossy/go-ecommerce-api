package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a JWT token with user information and expiration time
func GenerateToken(userID string, userRole string) (string, error) {
	// Define the expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(time.Hour * 1)

	// Create JWT claims with userID, role, and expiration time
	claims := &jwt.MapClaims{
		"sub":  userID,
		"role": userRole,
		"exp":  expirationTime.Unix(),
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken checks the validity of the provided JWT token.
func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userID := claims["sub"].(string)
	return userID, nil
}
