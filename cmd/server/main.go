package main

import (
	"log"
	"net/http"

	"Alya-Ecommerce-Go/internal/routers"

	"github.com/joho/godotenv"
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

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
