package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	util "Alya-Ecommerce-Go/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) InsertUser(ctx *fiber.Ctx) error {
	var user dto.InsertUserRequest
	err := ctx.BodyParser(&user)

	if err != nil {
		return util.GenerateResponse(ctx, http.StatusBadGateway, "Invalid Request", err.Error())
	}

	if errorMessage := util.ValidateData(&user); len(errorMessage) > 0 {
		return util.GenerateResponse(ctx, http.StatusBadGateway, "Validation Error, Please More Carefully Insert The Data", errorMessage)
	}

	hashPassword, _ := util.HashPassword(user.Password)

	data, _, err := c.Client.From("users").Insert(map[string]interface{}{
		"username":     user.Username,
		"password":     string(hashPassword),
		"full_name":    user.Name,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"address":      user.Address,
		"created_by":   user.Username,
		"created_at":   time.Now(),
	}, false, "", "", "").Execute()

	if err != nil {
		return util.GenerateResponse(ctx, http.StatusInternalServerError, "Failed to insert user", err.Error())
	}

	if data == nil {
		return util.GenerateResponse(ctx, http.StatusInternalServerError, "Success data not returned", "")
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success insert user", "")
}
