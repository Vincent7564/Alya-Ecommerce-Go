package routers

import (
	"log"

	"Alya-Ecommerce-Go/internal/controllers"
	user_service "Alya-Ecommerce-Go/internal/services"
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
	userService := user_service.NewUserService(client)
	// Pass the Supabase client to the controller
	userController := controllers.NewUserController(userService)

	// Define your routes and handlers
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/register-users", userController.RegisterUsers).Methods("POST")
	router.HandleFunc("/auth-login", userController.AuthLogin).Methods("POST")
	return router
}
