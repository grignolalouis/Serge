package utils

import (
	"fmt"
)

// AppError represents an application error with a code and message
type AppError struct {
	Code    int
	Message string
	Err     error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewAppError creates a new app error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NotFound creates a 404 error
func NotFound(message string, err error) *AppError {
	return NewAppError(404, message, err)
}

// BadRequest creates a 400 error
func BadRequest(message string, err error) *AppError {
	return NewAppError(400, message, err)
}

// InternalError creates a 500 error
func InternalError(message string, err error) *AppError {
	return NewAppError(500, message, err)
}

// Unauthorized creates a 401 error
func Unauthorized(message string, err error) *AppError {
	return NewAppError(401, message, err)
}

// Forbidden creates a 403 error
func Forbidden(message string, err error) *AppError {
	return NewAppError(403, message, err)
}
