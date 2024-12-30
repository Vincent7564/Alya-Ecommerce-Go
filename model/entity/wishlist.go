package entity

import "time"

type Wishlist struct {
	ID        int       `json:"id"`
	UsersID   int       `json:"users_id" validate:"required"`
	ProductID int8      `json:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
