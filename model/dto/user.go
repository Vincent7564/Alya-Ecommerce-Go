package dto

type CheckPasswordRequest struct {
	Password string `json:"password"`
	UsersId  int    `json:"users_id"`
}

type UpdateProfileRequest struct {
	Username    string `json:"username" validate:"omitempty"`
	Password    string `json:"password" validate:"omitempty,gte=8,lte=30"`
	FullName    string `json:"full_name" validate:"omitempty"`
	Email       string `json:"email" validate:"omitempty"`
	Address     string `json:"address" validate:"omitempty"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,gte=10,lte=14"`
}
