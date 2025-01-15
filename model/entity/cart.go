package entity

import "time"

type Cart struct {
	ID        int       `json:"id" db:"id" validate:"omitempty"`
	UserID    int       `json:"user_id" db:"user_id" validate:"required"`
	ProductId int       `json:"product_id" db:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validate:"omitempty"`
	CreatedBy string    `json:"created_by" db:"created_by" validate:"omitempty"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validate:"omitempty"`
	UpdatedBy string    `json:"updated_by" db:"updated_by" validate:"omitempty"`
	Qty       int       `json:"qty" db:"qty" validate:"required"`
}
