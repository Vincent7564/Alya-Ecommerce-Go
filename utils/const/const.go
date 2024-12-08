package cons

import (
	"github.com/gofiber/fiber/v2"
)

// Error Message
var (
	ErrInvalidRequest      = fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	ErrDataExisted         = "Data Existed"
	ErrEmailExisted        = "Email Existed"
	ErrUsernameExisted     = "Username Existed"
	ErrValidationError     = "Validation Error"
	ErrSuccess             = "Success"
	ErrFailed              = "Failed"
	ErrAccountNotFound     = "Account not found"
	ErrIncorrectPassword   = "Incorrect Password"
	ErrLoginSuccess        = "Login Success"
	ErrInternalServerError = "Internal server error"
	ErrTokenExpired        = "Token Expired"
	ErrNotFound            = "Not Found"
)
