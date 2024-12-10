package dto

type AddProductRequest struct {
	ProductName       string `json:"product_name" validate:"required"`
	ProductPrice      int    `json:"product_price" validate:"required"`
	ProductStock      int    `json:"product_stock" validate:"required"`
	ProductCategoryId int    `json:"product_category_id" validate:"required"`
	Discount          int    `json:"discount" validate:"required"`
	Description       string `json:"description" validate:"required,gte=10"`
}
