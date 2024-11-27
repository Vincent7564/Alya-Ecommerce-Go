package dto

import "time"

type InsertUserRequest struct {
	ID          uint   `json:"id" validate:"omitempty"`
	Username    string `json:"username" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"gte=8,lte=30"`
	Email       string `json:"email" validate:"required"`
	Role        string `json:"role" validate:"omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
