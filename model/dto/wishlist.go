package dto

type AddWishlistRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	UsersID   int   `json:"users_id" validate:"required"`
}
