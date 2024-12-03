package entity

import "time"

type UsersToken struct {
	ID        int       `json:"id" db:"id"`
	UsersID   int       `json:"users_id" db:"id"`
	Token     string    `json:"token" db:"token"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}
