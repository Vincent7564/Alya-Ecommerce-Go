package entity

import "time"

type UserEntity struct {
	ID          int       `json:"id" db:"id" validate:"omitempty"`
	Username    string    `json:"username" db:"username" validate:"required"`
	Password    string    `json:"password" db:"password" validate:"required,gte=8"`
	FullName    string    `json:"full_name" db:"full_name" validate:"omitempty"`
	Email       string    `json:"email" db:"email" validate:"required,email"`
	Address     string    `json:"address" db:"address" validate:"omitempty"`
	PhoneNumber string    `json:"phone_number" db:"phone_number" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" validate:"omitempty"`
	CreatedBy   string    `json:"created_by" db:"created_by" validate:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" validate:"omitempty"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by" validate:"omitempty"`
}
