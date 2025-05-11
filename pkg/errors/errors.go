package errors

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

// RespondWithError sends an error response to the client
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
