package user_service

import (
	"encoding/json"
	"fmt"

	"github.com/supabase-community/supabase-go"
)

type UserService struct {
	Client *supabase.Client
}

func NewUserService(client *supabase.Client) *UserService {
	return &UserService{Client: client}
}

type Token struct {
	Token     string `json:"token"`
	IsActive  bool   `json:"is_active"`
	ID        int8   `json:"id"`
	UsersID   int8   `json:"users_id"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
}

func (us *UserService) ValidateToken(token string) (bool, error) {
	var datas []Token

	data, _, err := us.Client.From("users_token").
		Select("*", "exact", false).
		Eq("token", token).
		Eq("is_active", "TRUE").Execute()

	if err != nil {
		fmt.Printf("Error fetching token data: %v\n", err)
		return false, err
	}

	fmt.Printf("Result: %+v\n", datas)
	if err := json.Unmarshal(data, &datas); err != nil {
		fmt.Printf("Error unmarshalling token data: %v\n", err)
		return false, err
	}

	if len(datas) == 0 {
		return false, nil
	}

	return true, nil
}
