package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type PaginateData struct {
	Limit int32
	Page  int32
	Total int32
}

type Response struct {
	Message  string        `json:"message"`
	Data     interface{}   `json:"data"`
	Paginate *PaginateData `json:"paginate,omitempty"`
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

var validate *validator.Validate

func GenerateResponse(ctx *fiber.Ctx, statusCode int, respmsg string, result interface{}) error {
	resp := Response{
		Message:  respmsg,
		Data:     result,
		Paginate: nil,
	}
	return ctx.Status(statusCode).JSON(resp)

}

func GenerateResponsePaginate(ctx *fiber.Ctx, statusCode int, respmsg string, result interface{}, paginate PaginateData) error {
	resp := Response{
		Message:  respmsg,
		Data:     result,
		Paginate: &paginate,
	}
	return ctx.Status(statusCode).JSON(resp)

}

func Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	validate = validator.New()
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ValidateData(data interface{}) []string {
	var errorMessage []string
	listError := Validate(data)
	if len(listError) > 0 && listError[0].Error {
		for _, err := range listError {
			errorMessage = append(errorMessage, fmt.Sprintf(
				"%s: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
	}
	return errorMessage
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

func SendEmail(to string, subject string, content string) error {
	m := gomail.NewMessage()
	email := os.Getenv("USER_EMAIL")
	port := 587
	password := os.Getenv("USER_PASSWORD")
	if email == "" || password == "" {
		return fmt.Errorf("email credentials not set in environment variables")
	}
	m.SetHeader("From", email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer("smtp.gmail.com", port, email, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
