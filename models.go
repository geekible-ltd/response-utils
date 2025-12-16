package responseutils

import "fmt"

// Response represents a standard API response
// @Description Standard API response structure
type Response struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Message string      `json:"message,omitempty" example:"Operation completed successfully"`
}

// SuccessResponseDTO represents a successful API response
// @Description Successful API response structure
type SuccessResponseDTO struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty" example:"Operation completed successfully"`
}

// ErrorResponseDTO represents an error API response
// @Description Error API response structure
type ErrorResponseDTO struct {
	Success bool        `json:"success" example:"false"`
	Error   ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
// @Description Error detail structure
type ErrorDetail struct {
	Code    string      `json:"code" example:"ERR_001"`
	Message string      `json:"message" example:"An error occurred"`
	Details interface{} `json:"details,omitempty"`
}

// CreatedResponseDTO represents a 201 Created response
// @Description Created response structure
type CreatedResponseDTO struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message" example:"Resource created successfully"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// AppError represents an application-specific error
type ResponseError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	StatusCode int                    `json:"-"`
	Details    map[string]interface{} `json:"details,omitempty"`
}

// Error implements the error interface
func (e *ResponseError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
