package responseutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, err error) {
	if appErr, ok := err.(*ResponseError); ok {
		c.JSON(appErr.StatusCode, Response{
			Success: false,
			Error: map[string]interface{}{
				"code":    appErr.Code,
				"message": appErr.Message,
				"details": appErr.Details,
			},
		})
		return
	}

	// Default to internal server error for unknown errors
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error: map[string]interface{}{
			"code":    ErrCodeInternalServer,
			"message": "An unexpected error occurred",
			"details": map[string]interface{}{"error": err.Error()},
		},
	})
}

// CreatedResponse sends a 201 Created response
func CreatedResponse(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusCreated, data, message)
}

// CreatedResponse sends a 201 Created response
func UpdatedResponse(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusAccepted, data, message)
}

// OKResponse sends a 200 OK response
func OKResponse(c *gin.Context, data interface{}, message string) {
	SuccessResponse(c, http.StatusOK, data, message)
}

// NoContentResponse sends a 204 No Content response
func NoContentResponse(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// ListResponseWithPagination sends a paginated list response
func ListResponseWithPagination(c *gin.Context, data interface{}, pagination *Pagination) {
	c.JSON(http.StatusOK, ListResponse{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	})
}

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, pageSize, total int) *Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}

	return &Pagination{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}
