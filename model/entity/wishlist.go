package entity

import "time"

type Wishlist struct {
	ID        int       `json:"id"`
	UsersID   int       `json:"users_id" validate:"required"`
	ProductID int8      `json:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ProductName       string `json:"product_name"`
	ProductPrice      int    `json:"product_price"`
	ProductCategoryID int    `json:"product_category_id"`
}

type GetWishlist struct {
	ID        int       `json:"id"`
	UsersID   int       `json:"users_id" validate:"required"`
	ProductID int8      `json:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"products"`
}
