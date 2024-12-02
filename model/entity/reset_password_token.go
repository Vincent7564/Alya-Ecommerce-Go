package entity

import "time"

type ResetPasswordToken struct {
	Username           string    `json:"username" validate:"omitempty"`
	Email              string    `json:"email" validate:"required"`
	ResetPasswordToken string    `json:"reset_password_token" validate:"required"`
	CreatedAt          time.Time `json:"created_at"`
	ExpiredAt          time.Time `json:"expired_at"`
}
