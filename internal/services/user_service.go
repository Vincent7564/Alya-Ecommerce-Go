package user_service

import (
	"github.com/supabase-community/supabase-go"
)

type UserService struct {
	Client *supabase.Client
}

func NewUserService(client *supabase.Client) *UserService {
	return &UserService{Client: client}
}

func (us *UserService) ValidateToken(token string) (bool, error) {
	data := us.Client.From("user_token").
		Select("*", "exact", false).
		Eq("token", token).
		Eq("is_active", "1").Single()

	return data != nil, nil
}
