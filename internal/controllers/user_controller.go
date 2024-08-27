package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/supabase-community/supabase-go"
)

type UserController struct {
	Client *supabase.Client
}

func NewUserController(client *supabase.Client) *UserController {
	return &UserController{Client: client}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch data from Supabase
	data, _, err := uc.Client.From("users").Select("*", "exact", false).Execute()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Respond with the fetched data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
