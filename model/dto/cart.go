package dto

type AddCartRequest struct {
	UsersID   int `json:"users_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required"`
}

type GetCartResponse struct {
	UsersID int     `json:"user_id"`
	Qty     int     `json:"qty"`
	Product Product `json:"products"`
}

type Product struct {
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	ProductStock int    `json:"product_stock"`
}
