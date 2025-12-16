# Response Utils

A standardized HTTP response utility package for Go applications using the Gin web framework. This package provides consistent API response structures, error handling, and pagination support for building RESTful APIs.

## Features

- ✅ Standardized JSON response formats
- ✅ Pre-built success and error response functions
- ✅ Comprehensive error types with HTTP status codes
- ✅ Built-in pagination support
- ✅ Type-safe error handling
- ✅ Swagger/OpenAPI compatible response models
- ✅ Fluent API for adding error details

## Installation

```bash
go get github.com/geekible-ltd/response-utils
```

## Quick Start

```go
import (
    responseutils "github.com/geekible-ltd/response-utils"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    r.GET("/users/:id", func(c *gin.Context) {
        user, err := getUserByID(c.Param("id"))
        if err != nil {
            responseutils.ErrorResponse(c, responseutils.NotFound("User"))
            return
        }
        responseutils.OKResponse(c, user, "User retrieved successfully")
    })
    
    r.Run(":8080")
}
```

## Response Structures

### Success Response

```json
{
    "success": true,
    "data": { ... },
    "message": "Operation completed successfully"
}
```

### Error Response

```json
{
    "success": false,
    "error": {
        "code": "NOT_FOUND",
        "message": "User not found",
        "details": { ... }
    }
}
```

### Paginated List Response

```json
{
    "success": true,
    "data": [ ... ],
    "pagination": {
        "page": 1,
        "page_size": 20,
        "total": 100,
        "total_pages": 5
    }
}
```

## Usage Examples

### 1. Success Responses

#### OK Response (200)

```go
func GetUser(c *gin.Context) {
    user := User{ID: "123", Name: "John Doe"}
    responseutils.OKResponse(c, user, "User retrieved successfully")
}
```

#### Created Response (201)

```go
func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        responseutils.ErrorResponse(c, responseutils.ValidationError("Invalid user data"))
        return
    }
    
    // Create user in database...
    
    responseutils.CreatedResponse(c, user, "User created successfully")
}
```

#### Updated Response (202)

```go
func UpdateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        responseutils.ErrorResponse(c, responseutils.ValidationError("Invalid user data"))
        return
    }
    
    // Update user in database...
    
    responseutils.UpdatedResponse(c, user, "User updated successfully")
}
```

#### No Content Response (204)

```go
func DeleteUser(c *gin.Context) {
    // Delete user from database...
    
    responseutils.NoContentResponse(c)
}
```

### 2. Error Responses

#### Using Pre-defined Error Functions

```go
func GetUser(c *gin.Context) {
    id := c.Param("id")
    
    user, err := userService.GetByID(id)
    if err != nil {
        // Returns 404 with standard error format
        responseutils.ErrorResponse(c, responseutils.NotFound("User"))
        return
    }
    
    responseutils.OKResponse(c, user, "User retrieved successfully")
}
```

#### Available Error Functions

```go
// Bad Request (400)
responseutils.BadRequest("Invalid request parameters")

// Unauthorized (401)
responseutils.Unauthorized("Invalid credentials")

// Forbidden (403)
responseutils.Forbidden("Access denied")

// Not Found (404)
responseutils.NotFound("Resource")

// Conflict (409)
responseutils.Conflict("Resource already exists")

// Validation Error (400)
responseutils.ValidationError("Validation failed")

// Internal Server Error (500)
responseutils.InternalServerError("Something went wrong")

// Database Error (500)
responseutils.DatabaseError(err)

// Invalid Input (400)
responseutils.InvalidInput("email", "must be a valid email address")

// Missing Header (400)
responseutils.MissingHeader("X-API-Key")

// Invalid UUID (400)
responseutils.InvalidUUID("user_id")

// Duplicate Entry (409)
responseutils.DuplicateEntry("User")

// Foreign Key Violation (400)
responseutils.ForeignKeyViolation("Referenced entity does not exist")
```

#### Creating Custom Errors

```go
func CustomErrorExample(c *gin.Context) {
    err := responseutils.NewResponseError(
        "CUSTOM_ERROR",
        "This is a custom error message",
        http.StatusTeapot, // 418
    )
    
    responseutils.ErrorResponse(c, err)
}
```

#### Adding Error Details

```go
func ValidateUser(c *gin.Context) {
    err := responseutils.ValidationError("User validation failed").
        WithDetails("email", "Email is required").
        WithDetails("age", "Must be at least 18 years old")
    
    responseutils.ErrorResponse(c, err)
}

// Response:
// {
//     "success": false,
//     "error": {
//         "code": "VALIDATION_ERROR",
//         "message": "User validation failed",
//         "details": {
//             "email": "Email is required",
//             "age": "Must be at least 18 years old"
//         }
//     }
// }
```

### 3. Pagination

#### Simple Pagination

```go
func ListUsers(c *gin.Context) {
    page := 1
    pageSize := 20
    
    // Get page and pageSize from query params
    if p := c.Query("page"); p != "" {
        page, _ = strconv.Atoi(p)
    }
    if ps := c.Query("page_size"); ps != "" {
        pageSize, _ = strconv.Atoi(ps)
    }
    
    // Fetch users from database
    users, total, err := userService.List(page, pageSize)
    if err != nil {
        responseutils.ErrorResponse(c, responseutils.DatabaseError(err))
        return
    }
    
    // Calculate pagination metadata
    pagination := responseutils.CalculatePagination(page, pageSize, total)
    
    // Send paginated response
    responseutils.ListResponseWithPagination(c, users, pagination)
}

// Response:
// {
//     "success": true,
//     "data": [
//         { "id": "1", "name": "John" },
//         { "id": "2", "name": "Jane" }
//     ],
//     "pagination": {
//         "page": 1,
//         "page_size": 20,
//         "total": 100,
//         "total_pages": 5
//     }
// }
```

#### Manual Pagination

```go
func ListProducts(c *gin.Context) {
    products := getProducts()
    
    pagination := &responseutils.Pagination{
        Page:       1,
        PageSize:   10,
        Total:      250,
        TotalPages: 25,
    }
    
    responseutils.ListResponseWithPagination(c, products, pagination)
}
```

## Error Codes Reference

| Error Code | HTTP Status | Description |
|------------|-------------|-------------|
| `BAD_REQUEST` | 400 | Generic bad request |
| `UNAUTHORIZED` | 401 | Authentication required |
| `FORBIDDEN` | 403 | Access denied |
| `NOT_FOUND` | 404 | Resource not found |
| `CONFLICT` | 409 | Resource conflict |
| `VALIDATION_ERROR` | 400 | Input validation failed |
| `INTERNAL_SERVER_ERROR` | 500 | Server error |
| `DATABASE_ERROR` | 500 | Database operation failed |
| `INVALID_INPUT` | 400 | Invalid field input |
| `MISSING_HEADER` | 400 | Required header missing |
| `INVALID_UUID` | 400 | Invalid UUID format |
| `DUPLICATE_ENTRY` | 409 | Duplicate resource |
| `FOREIGN_KEY_VIOLATION` | 400 | Foreign key constraint violation |
| `ACCOUNT_LOCKED` | 403 | User account locked |

## API Reference

### Response Functions

#### `OKResponse(c *gin.Context, data interface{}, message string)`
Sends a 200 OK response with data and message.

#### `CreatedResponse(c *gin.Context, data interface{}, message string)`
Sends a 201 Created response with data and message.

#### `UpdatedResponse(c *gin.Context, data interface{}, message string)`
Sends a 202 Accepted response with data and message.

#### `NoContentResponse(c *gin.Context)`
Sends a 204 No Content response (typically for DELETE operations).

#### `ErrorResponse(c *gin.Context, err error)`
Sends an error response. Automatically handles `*ResponseError` types with proper status codes and formatting.

#### `ListResponseWithPagination(c *gin.Context, data interface{}, pagination *Pagination)`
Sends a paginated list response with data and pagination metadata.

#### `SuccessResponse(c *gin.Context, statusCode int, data interface{}, message string)`
Generic success response function with custom status code.

### Pagination Functions

#### `CalculatePagination(page, pageSize, total int) *Pagination`
Calculates pagination metadata from page number, page size, and total count.

- Defaults: `page = 1`, `pageSize = 20` if invalid values provided
- Returns: `*Pagination` with calculated `TotalPages`

### Error Creation Functions

All error functions return `*ResponseError` which implements the `error` interface.

#### `NewResponseError(code string, message string, statusCode int) *ResponseError`
Creates a custom error with specified code, message, and HTTP status code.

#### `WithDetails(key string, value interface{}) *ResponseError`
Chains additional details to an error. Returns `*ResponseError` for fluent chaining.

## Complete Example

Here's a complete example of a simple CRUD API:

```go
package main

import (
    "github.com/gin-gonic/gin"
    responseutils "github.com/geekible-ltd/response-utils"
    "strconv"
)

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = make(map[string]User)

func main() {
    r := gin.Default()
    
    // List users with pagination
    r.GET("/users", func(c *gin.Context) {
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
        
        userList := make([]User, 0, len(users))
        for _, user := range users {
            userList = append(userList, user)
        }
        
        pagination := responseutils.CalculatePagination(page, pageSize, len(userList))
        responseutils.ListResponseWithPagination(c, userList, pagination)
    })
    
    // Get user by ID
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        user, exists := users[id]
        
        if !exists {
            responseutils.ErrorResponse(c, responseutils.NotFound("User"))
            return
        }
        
        responseutils.OKResponse(c, user, "User retrieved successfully")
    })
    
    // Create user
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            responseutils.ErrorResponse(c, 
                responseutils.ValidationError("Invalid user data").
                    WithDetails("error", err.Error()))
            return
        }
        
        if _, exists := users[user.ID]; exists {
            responseutils.ErrorResponse(c, responseutils.DuplicateEntry("User"))
            return
        }
        
        users[user.ID] = user
        responseutils.CreatedResponse(c, user, "User created successfully")
    })
    
    // Update user
    r.PUT("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        
        if _, exists := users[id]; !exists {
            responseutils.ErrorResponse(c, responseutils.NotFound("User"))
            return
        }
        
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            responseutils.ErrorResponse(c, 
                responseutils.ValidationError("Invalid user data"))
            return
        }
        
        user.ID = id
        users[id] = user
        responseutils.UpdatedResponse(c, user, "User updated successfully")
    })
    
    // Delete user
    r.DELETE("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        
        if _, exists := users[id]; !exists {
            responseutils.ErrorResponse(c, responseutils.NotFound("User"))
            return
        }
        
        delete(users, id)
        responseutils.NoContentResponse(c)
    })
    
    r.Run(":8080")
}
```

## Best Practices

1. **Consistent Error Handling**: Always use `ErrorResponse()` for errors to maintain consistent error format across your API.

2. **Meaningful Messages**: Provide clear, actionable error messages and success messages.

3. **Use Appropriate HTTP Status Codes**: Use the pre-built response functions that set appropriate status codes.

4. **Add Error Details**: Use `WithDetails()` to provide additional context for debugging:
   ```go
   err := responseutils.ValidationError("Invalid data").
       WithDetails("field", "email").
       WithDetails("reason", "must be unique")
   ```

5. **Pagination Defaults**: Use `CalculatePagination()` for consistent pagination behavior with automatic defaults.

6. **Custom Errors**: For domain-specific errors, create custom error codes:
   ```go
   const (
       ErrCodeSubscriptionExpired = "SUBSCRIPTION_EXPIRED"
       ErrCodeQuotaExceeded = "QUOTA_EXCEEDED"
   )
   
   func SubscriptionExpired() *responseutils.ResponseError {
       return responseutils.NewResponseError(
           ErrCodeSubscriptionExpired,
           "Your subscription has expired",
           http.StatusPaymentRequired,
       )
   }
   ```

## Requirements

- Go 1.24.5 or higher
- [Gin Web Framework](https://github.com/gin-gonic/gin) v1.11.0 or higher

## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues, questions, or contributions, please open an issue in the repository.

