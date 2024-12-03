package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) CheckPassword(ctx *fiber.Ctx) error {
	var request dto.CheckPasswordRequest
	var user entity.UserEntity
	err := ctx.BodyParser(&request)

	if err != nil {
		return util.GenerateResponse(ctx, http.StatusBadGateway, cons.ErrInvalidRequest, "")
	}

	_, err = c.Client.From("users").Select("*", "", false).Eq("id", string(request.UsersId)).Single().ExecuteTo(&user)

	if err != nil {
		return util.GenerateResponse(ctx, http.StatusBadGateway, cons.ErrAccountNotFound, "")
	}

	isTrue := util.CheckPasswordHash(request.Password, user.Password)

	if isTrue {
		return util.GenerateResponse(ctx, http.StatusOK, "Password Correct", "")
	} else {
		return util.GenerateResponse(ctx, http.StatusBadGateway, "Incorrect Password", "")
	}
}
