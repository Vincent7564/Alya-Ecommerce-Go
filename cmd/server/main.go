package main

import (
	"log"
	"net/http"

	"Alya-Ecommerce-Go/internal/routers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnv()

	// Setup router
	router := routers.SetupRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your Next.js frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
