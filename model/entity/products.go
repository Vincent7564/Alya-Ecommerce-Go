package entity

type Products struct {
	ID                int64  `json:"id"`
	ProductName       string `json:"product_name"`
	ProductStock      int64  `json:"product_stock"`
	ProductPrice      int64  `json:"product_price"`
	ProductCategoryID int64  `json:"product_category_id"`
	Category          struct {
		CategoryName string `json:"category_name"`
	} `json:"category"`
	ProductDescription string `json:"description"`
	Discount           int64  `json:"discount"`
	CategoryName       string `json:"category_name"`
}
