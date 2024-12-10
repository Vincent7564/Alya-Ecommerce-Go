package util

import (
	"Alya-Ecommerce-Go/model/entity"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/supabase-community/supabase-go"
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

func CheckToken(c *supabase.Client, token string) bool {
	var entityTokens entity.UsersToken

	claims, err := DecodeToken(token)

	if err != nil {
		return false
	}

	exp, ok := (*claims)["exp"].(float64)

	if !ok {
		return false
	}

	expirationTime := time.Unix(int64(exp), 0)
	if time.Now().After(expirationTime) {
		return false
	}

	counter, err := c.From("users_token").Select("*", "", false).Eq("token", token).Single().ExecuteTo(&entityTokens)

	if err != nil {
		return false
	}

	if counter != 0 {
		return false
	}

	if time.Now().After(entityTokens.ExpiresAt) {
		return false
	}
	return true
}

func DecodeToken(tokenString string) (*jwt.MapClaims, error) {
	var secret_token = os.Getenv("SECRET_TOKEN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret_token), nil
	})

	if err != nil {
		fmt.Printf("error Parsing Token " + err.Error())
		return nil, fmt.Errorf("error Parsing Token :%w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid Token")
}
