package responseutils

import (
	"fmt"
	"net/http"
)

// Error codes
const (
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeForbidden           = "FORBIDDEN"
	ErrCodeNotFound            = "NOT_FOUND"
	ErrCodeConflict            = "CONFLICT"
	ErrCodeValidation          = "VALIDATION_ERROR"
	ErrCodeInternalServer      = "INTERNAL_SERVER_ERROR"
	ErrCodeDatabase            = "DATABASE_ERROR"
	ErrCodeInvalidInput        = "INVALID_INPUT"
	ErrCodeMissingHeader       = "MISSING_HEADER"
	ErrCodeInvalidUUID         = "INVALID_UUID"
	ErrCodeDuplicateEntry      = "DUPLICATE_ENTRY"
	ErrCodeForeignKeyViolation = "FOREIGN_KEY_VIOLATION"
	ErrCodeInvalidBody         = "INVALID_BODY"
	ErrUserAccountLocked       = "ACCOUNT_LOCKED"
	ErrUnauthorizedError       = "UNAUTHORIZED_ERROR"
)

// NewResponseError creates a new ResponseError
func NewResponseError(code string, message string, statusCode int) *ResponseError {
	return &ResponseError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}
}

// WithDetails adds details to the error
func (e *ResponseError) WithDetails(key string, value interface{}) *ResponseError {
	e.Details[key] = value
	return e
}

// Common errors
func BadRequest(message string) *ResponseError {
	return NewResponseError(ErrCodeBadRequest, message, http.StatusBadRequest)
}

func Unauthorized(message string) *ResponseError {
	return NewResponseError(ErrCodeUnauthorized, message, http.StatusUnauthorized)
}

func Forbidden(message string) *ResponseError {
	return NewResponseError(ErrCodeForbidden, message, http.StatusForbidden)
}

func NotFound(resource string) *ResponseError {
	return NewResponseError(
		ErrCodeNotFound,
		fmt.Sprintf("%s not found", resource),
		http.StatusNotFound,
	)
}

func Conflict(message string) *ResponseError {
	return NewResponseError(ErrCodeConflict, message, http.StatusConflict)
}

func ValidationError(message string) *ResponseError {
	return NewResponseError(ErrCodeValidation, message, http.StatusBadRequest)
}

func InternalServerError(message string) *ResponseError {
	return NewResponseError(ErrCodeInternalServer, message, http.StatusInternalServerError)
}

func DatabaseError(err error) *ResponseError {
	return NewResponseError(
		ErrCodeDatabase,
		"Database operation failed",
		http.StatusInternalServerError,
	).WithDetails("error", err.Error())
}

func InvalidInput(field string, reason string) *ResponseError {
	return NewResponseError(
		ErrCodeInvalidInput,
		fmt.Sprintf("Invalid input for field '%s': %s", field, reason),
		http.StatusBadRequest,
	)
}

func MissingHeader(header string) *ResponseError {
	return NewResponseError(
		ErrCodeMissingHeader,
		fmt.Sprintf("Missing required header: %s", header),
		http.StatusBadRequest,
	)
}

func InvalidUUID(field string) *ResponseError {
	return NewResponseError(
		ErrCodeInvalidUUID,
		fmt.Sprintf("Invalid UUID format for field '%s'", field),
		http.StatusBadRequest,
	)
}

func DuplicateEntry(resource string) *ResponseError {
	return NewResponseError(
		ErrCodeDuplicateEntry,
		fmt.Sprintf("%s already exists", resource),
		http.StatusConflict,
	)
}

func ForeignKeyViolation(message string) *ResponseError {
	return NewResponseError(
		ErrCodeForeignKeyViolation,
		message,
		http.StatusBadRequest,
	)
}

func UnauthorizedError(message string) *ResponseError {
	return NewResponseError(
		ErrCodeUnauthorized,
		message,
		http.StatusUnauthorized,
	)
}

func VersionExistsError(versionNumber string) *ResponseError {
	return NewResponseError(
		ErrCodeConflict,
		fmt.Sprintf("Version %s already exists", versionNumber),
		http.StatusConflict,
	)
}
