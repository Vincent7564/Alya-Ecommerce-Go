package entity

import "time"

type Categories struct {
	CategoryName string    `json:"category_name"`
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
}
