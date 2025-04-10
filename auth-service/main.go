package main

import (
	"Alya-Ecommerce-Go/auth-service/routes"
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
	"go.opentelemetry.io/otel"
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

	log.Info().Msg("This is an info message")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")

	log.Info().
		Str("module", "main").
		Int("status", 200).
		Msg("Operation completed successfully")

	// tp, err := tracing.InitializeTracerProvider("alya-ecomm/auth-service") // Use the correct service name
	// if err != nil {
	// 	log.Error().Msg("Failed to initialize tracer provider " + err.Error())
	// }
	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		log.Error().Msg("Error shutting down tracer provider " + err.Error())
	// 	}
	// }()

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandling,
	})
	app.Use(cors.New())
	// app.Use(func(c *fiber.Ctx) error {
	// 	tracer := otel.Tracer("http-request")
	// 	ctx, span := tracer.Start(c.UserContext(), c.Method()+" "+c.Path())
	// 	defer span.End()
	// 	c.SetUserContext(ctx)
	// 	return c.Next()
	// })
	SupabaseURL := os.Getenv("NEXT_PUBLIC_SUPABASE_URL")
	SupabaseAnon := os.Getenv("NEXT_PUBLIC_SUPABASE_ANON")
	client, err := supabase.NewClient(SupabaseURL, SupabaseAnon, &supabase.ClientOptions{})
	if err != nil {
		fmt.Printf("Cannot connect to Supabase!")
	}
	routes.InitRouter(app, client)
	app.Get("/", func(c *fiber.Ctx) error {
		tracer := otel.Tracer("http-handler")

		_, span := tracer.Start(c.UserContext(), "handle-root-request")

		defer span.End()

		return c.SendString("Hello, World!")
	})
	app.Listen(":8020")
}
