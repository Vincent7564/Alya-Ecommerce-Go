package dto

type AddProductRequest struct {
	ProductName       string `json:"product_name" validate:"required,gte=10"`
	ProductPrice      int    `json:"product_price" validate:"required"`
	ProductStock      int    `json:"product_stock" validate:"required"`
	ProductCategoryId int    `json:"product_category_id" validate:"required"`
	Discount          int    `json:"discount"`
	Description       string `json:"description" validate:"required,gte=10"`
}

type AddCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type UpdateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type UpdateProductRequest struct {
	ProductName       string `json:"product_name" validate:"required,gte=10"`
	ProductPrice      int    `json:"product_price" validate:"required"`
	ProductStock      int    `json:"product_stock" validate:"required"`
	ProductCategoryId int    `json:"product_category_id" validate:"required"`
	Discount          int    `json:"discount"`
	Description       string `json:"description" validate:"required,gte=10"`
}
