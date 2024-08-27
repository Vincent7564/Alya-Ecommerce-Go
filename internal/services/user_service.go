package services

import (
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

func GetUserByID(id string) (models.User, error) {
	return repositories.FindUserByID(id)
}
