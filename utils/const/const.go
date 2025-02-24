package cons

import (
	"github.com/gofiber/fiber/v2"
)

// Error Message
var (
	ErrInvalidRequest      = fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
	ErrTokenFailed         = fiber.NewError(fiber.StatusBadRequest, "Failed")
	ErrDataExisted         = fiber.NewError(fiber.StatusBadGateway, "Data Existed")
	ErrEmailExisted        = fiber.NewError(fiber.StatusBadGateway, "Email Existed")
	ErrUsernameExisted     = fiber.NewError(fiber.StatusBadGateway, "Username Existed")
	ErrValidationError     = fiber.NewError(fiber.StatusBadGateway, "Validation Error")
	ErrSuccess             = fiber.NewError(fiber.StatusOK, "Success")
	ErrFailed              = "Failed"
	ErrAccountNotFound     = fiber.NewError(fiber.StatusNotFound, "Account not found")
	ErrIncorrectPassword   = fiber.NewError(fiber.StatusBadGateway, "Invalid password")
	ErrLoginSuccess        = "Login Success"
	ErrInternalServerError = fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	ErrTokenExpired        = "Token Expired"
	ErrNotFound            = "Not Found"
	ErrDataNotFound        = fiber.NewError(fiber.StatusBadGateway, "Data not Found")
)
