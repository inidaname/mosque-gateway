package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
type ResponseWithMeta struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	success := statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
	response := Response{
		Success: success,
		Message: message,
		Data:    data,
	}

	// Set the HTTP status code
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func SendResponseWithMeta(w http.ResponseWriter, statusCode int, message string, data interface{}, meta interface{}) {
	// Determine success based on status code
	success := statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices

	response := ResponseWithMeta{
		Success: success,
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	// Set the HTTP status code
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	err := json.NewDecoder(r.Body).Decode(data)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err != nil {
		var syntaxError *json.SyntaxError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("badly-formed JSON at character %d", syntaxError.Offset)
		default:
			return err
		}
	}
	return nil
}

func ReadIDParam(r *http.Request, inputID string) (uuid.UUID, error) {
	idStr := chi.URLParam(r, inputID)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.New(), errors.New("id must be a valid UUID")
	}
	return id, nil
}

func ReadAccountNumberParam(r *http.Request, paramName string) (string, error) {
	accountNumber := chi.URLParam(r, paramName)
	if matched, _ := regexp.MatchString(`^\d{10}$`, accountNumber); !matched {
		return "", errors.New("account number must be exactly 10 digits")
	}

	return accountNumber, nil
}

var formDecoder = form.NewDecoder()

// ParseBody reads request body into the provided struct
func ParseBody(r *http.Request, dst interface{}) error {
	contentType := r.Header.Get("Content-Type")

	switch {
	case strings.Contains(contentType, "application/json"):
		return json.NewDecoder(r.Body).Decode(dst)

	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		if err := r.ParseForm(); err != nil {
			return err
		}
		return formDecoder.Decode(dst, r.Form)

	case strings.Contains(contentType, "multipart/form-data"):
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			return err
		}
		return formDecoder.Decode(dst, r.MultipartForm.Value)

	default:
		return errors.New("unsupported Content-Type")
	}
}

func SafeStringToPgText(strPtr *string) pgtype.Text {
	if strPtr != nil && *strPtr != "" {
		return pgtype.Text{String: *strPtr, Valid: true}
	}
	return pgtype.Text{Valid: false}
}
func Deref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
