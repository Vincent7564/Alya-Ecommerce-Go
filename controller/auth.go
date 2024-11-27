package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	util "Alya-Ecommerce-Go/utils"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var request dto.LoginRequest
	err := ctx.BodyParser(&request)

	var getData dto.InsertUserRequest

	if err != nil {
		return util.GenerateResponse(ctx, http.StatusBadGateway, "Invalid Request", err.Error())
	}
	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		return util.GenerateResponse(ctx, http.StatusBadGateway, "Validate Error", errorMessage)
	}

	count, err := c.Client.From("users").Select("*", "", false).Eq("username", request.Username).ExecuteTo(&getData)

	if err != nil && count == 0 {
		return util.GenerateResponse(ctx, http.StatusNotFound, "Account Not Found", err.Error())
	}

	isTrue := util.CheckPasswordHash(request.Password, getData.Password)

	if isTrue {
		claims := jwt.MapClaims{
			"username": getData.Username,
			"email":    getData.Email,
			"exp":      time.Now().Add(time.Hour * 3).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		secret_token := os.Getenv("SECRET_TOKEN")
		t, err := token.SignedString([]byte(secret_token))

		if err != nil {
			return util.GenerateResponse(ctx, http.StatusInternalServerError, "Failed to Sign Token", err)
		}

		data, _, err := c.Client.From("users_token").Insert(map[string]interface{}{
			"users_id":   getData.ID,
			"token":      t,
			"is_active":  true,
			"created_at": time.Now(),
			"expires_at": time.Now().Add(time.Hour * 12),
		}, false, "", "", "").Execute()

		if err != nil {
			return util.GenerateResponse(ctx, http.StatusBadGateway, "Failed to insert token", err)
		}

		if data == nil {
			return util.GenerateResponse(ctx, http.StatusInternalServerError, "No Data Returned", data)
		}

		return util.GenerateResponse(ctx, http.StatusOK, "Login Success", t)
	}
	return util.GenerateResponse(ctx, http.StatusBadGateway, "Login Failed", "")
}
