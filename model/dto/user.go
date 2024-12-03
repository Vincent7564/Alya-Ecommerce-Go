package dto

type CheckPasswordRequest struct {
	Password string `json:"password"`
	UsersId  int    `json:"users_id"`
}
