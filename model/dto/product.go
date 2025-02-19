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

type GetProductBySearch struct {
	ProductName string `json:"product_name"`
}

type RetrieveProductParameter struct {
	SearchQuery *string `json:"search_query,omitempty"`
	MinPrice    *int    `json:"min_price,omitempty"`
	MaxPrice    *int    `json:"max_price,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
	SortBy      *string `json:"sort_by,omitempty"`
	Order       *string `json:"order,omitempty"`
}
