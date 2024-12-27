package utils

import (
	"net/http"
	"encoding/json"
)

// RespondWithJSON is a utility function to send a JSON response with a specific status code.
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// RespondWithError is a utility function to send a JSON error response.
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}

// ValidateEmail checks if the provided email is valid.
func ValidateEmail(email string) bool {
	// Simple email validation logic (can be improved)
	if len(email) == 0 {
		return false
	}
	return true
}

// ValidatePassword checks if the provided password meets certain criteria.
func ValidatePassword(password string) bool {
	// Simple password validation logic (can be improved)
	if len(password) < 6 {
		return false
	}
	return true
}