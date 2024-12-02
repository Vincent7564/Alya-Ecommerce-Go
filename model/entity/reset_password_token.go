package entity

import "time"

type ResetPasswordToken struct {
	UserId             int8      `json:"users_id"`
	Email              string    `json:"email"`
	ResetPasswordToken string    `json:"reset_password_token"`
	CreatedAt          time.Time `json:"created_at"`
	ExpiredAt          time.Time `json:"expired_at"`
}
