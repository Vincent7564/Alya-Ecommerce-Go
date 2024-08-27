package routers

import (
	"log"

	"Alya-Ecommerce-Go/internal/controllers"
	"os"

	"github.com/gorilla/mux"
	"github.com/supabase-community/supabase-go"
)

func SetupRouter() *mux.Router {
	// Initialize the router
	router := mux.NewRouter()

	// Load environment variables
	URL := os.Getenv("NEXT_PUBLIC_SUPABASE_URL")
	ANON := os.Getenv("NEXT_PUBLIC_SUPABASE_ANON")

	// Create Supabase client
	client, err := supabase.NewClient(URL, ANON, &supabase.ClientOptions{})
	if err != nil {
		log.Fatalf("Failed to create Supabase client: %v", err)
	}

	// Pass the Supabase client to the controller
	userController := controllers.NewUserController(client)

	// Define your routes and handlers
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")

	return router
}
