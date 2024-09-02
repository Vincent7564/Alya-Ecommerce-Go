package controllers

import (
	user_service "Alya-Ecommerce-Go/internal/services"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service *user_service.UserService
}

func NewUserController(client *user_service.UserService) *UserController {
	return &UserController{Service: client}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, _, err := uc.Service.Client.From("users").Select("*", "exact", false).Execute()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (uc *UserController) RegisterUsers(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Fullname    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
	}
	type Response struct {
		StatusCode int
		Message    string
	}
	response := Response{
		StatusCode: 200,
		Message:    "Successs",
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request, Please Check", http.StatusBadRequest)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Encrypt Password Failed", http.StatusInternalServerError)
		return
	}

	t := time.Now()

	data, _, err := uc.Service.Client.From("users").Insert(map[string]interface{}{
		"username":     user.Username,
		"password":     string(hashPassword),
		"full_name":    user.Fullname,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"address":      user.Address,
		"created_by":   user.Username,
		"created_at":   t,
	}, false, "", "", "").Execute()

	if err != nil {
		http.Error(w, "Failed to Register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if data != nil {
		response.StatusCode = 200
		response.Message = "Success"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
	}

	json.NewEncoder(w).Encode(response)
}
