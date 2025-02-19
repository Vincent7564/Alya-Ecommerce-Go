package entity

type Category struct {
	CategoryName string `json:"category_name"`
}

type ProductImages struct {
	Images string `json:"image_url"`
}

type Products struct {
	ID                 int64           `json:"id"`
	ProductName        string          `json:"product_name"`
	ProductStock       int64           `json:"product_stock"`
	ProductPrice       int64           `json:"product_price"`
	ProductCategoryID  int64           `json:"product_category_id"`
	Discount           int64           `json:"discount"`
	ProductDescription string          `json:"description"`
	Created_by         string          `json:"created_by"`
	ProductImages      []ProductImages `json:"product_images"`
	Category           Category        `json:"category"`
}
