package main

import (
	"Alya-Ecommerce-Go/router"
	util "Alya-Ecommerce-Go/utils"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/supabase-community/supabase-go"
)

func ErrorHandling(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return util.GenerateResponse(ctx, code, err.Error(), "")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()

	// Example logs
	log.Info().Msg("This is an info message")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")

	// Log with additional fields
	log.Info().
		Str("module", "main").
		Int("status", 200).
		Msg("Operation completed successfully")

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandling,
	})
	app.Use(cors.New())
	SupabaseURL := os.Getenv("NEXT_PUBLIC_SUPABASE_URL")
	SupabaseAnon := os.Getenv("NEXT_PUBLIC_SUPABASE_ANON")
	client, err := supabase.NewClient(SupabaseURL, SupabaseAnon, &supabase.ClientOptions{})

	if err != nil {
		fmt.Printf("Cannot connect to Supabase!")
	}
	router.InitRouter(app, client)
	app.Listen("localhost:8080")
}
