package main

import (
	"Alya-Ecommerce-Go/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()
	app.Use(cors.New())
	SupabaseURL := os.Getenv("NEXT_PUBLIC_SUPABASE_URL")
	SupabaseAnon := os.Getenv("NEXT_PUBLIC_SUPABASE_ANON")
	client, err := supabase.NewClient(SupabaseURL, SupabaseAnon, &supabase.ClientOptions{})

	if err != nil {
		log.Fatalf("Cannot connect to Supabase!")
	}
	router.InitRouter(app, client)
	app.Listen(":8080")
}
