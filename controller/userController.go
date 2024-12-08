package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) CheckPassword(ctx *fiber.Ctx) error {
	var request dto.CheckPasswordRequest
	var user entity.UserEntity
	FuncName := "CheckPassword"
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	_, err = c.Client.From("users").Select("*", "", false).Eq("id", string(request.UsersId)).Single().ExecuteTo(&user)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrAccountNotFound
	}

	isTrue := util.CheckPasswordHash(request.Password, user.Password)

	if isTrue {
		return util.GenerateResponse(ctx, http.StatusOK, "Password Correct", "")
	} else {
		log.Error().Err(&fiber.Error{Message: "Incorrect Password"}).Msg("API Endpoint /" + FuncName)
		return cons.ErrIncorrectPassword
	}
}

func (c *Controller) UpdateProfile(ctx *fiber.Ctx) error {
	var request dto.UpdateProfileRequest
	FuncName := "UpdateProfile"
	idParam := ctx.Params("id")
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	hashedPassword, err := util.HashPassword(request.Password)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusInternalServerError, cons.ErrInternalServerError, "")
	}

	_, _, err = c.Client.From("users").Update(map[string]interface{}{
		"username":     request.Username,
		"password":     hashedPassword,
		"full_name":    request.FullName,
		"email":        request.Email,
		"address":      request.Address,
		"phone_number": request.PhoneNumber,
		"updated_at":   time.Now(),
		"updated_by":   request.Username,
	}, "", "").Eq("id", idParam).Single().Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusBadGateway, cons.ErrFailed+" to update ", err.Error())
	}

	return cons.ErrSuccess
}
