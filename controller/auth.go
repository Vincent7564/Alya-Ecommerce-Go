package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) InsertUser(ctx *fiber.Ctx) error {
	var user dto.InsertUserRequest
	FuncName := "InsertUser"
	err := ctx.BodyParser(&user)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	URL := "http://127.0.0.1:8020/auth/login"
	respCode, _ := util.HitMicroservicesAPI(FuncName, URL, "POST", "application/json", user)

	if respCode != 200 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", user.Email)
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var request dto.LoginRequest
	var response dto.LoginResponse
	FuncName := "Login"
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	URL := "http://127.0.0.1:8020/auth/login"
	respCode, resp := util.HitMicroservicesAPI(FuncName, URL, "POST", "application/json", request)

	if respCode != 200 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return util.GenerateResponse(ctx, http.StatusOK, "Success", response)
}

func (c *Controller) ForgotPassword(ctx *fiber.Ctx) error {
	var request dto.ForgotPasswordRequest
	FuncName := "ForgotPassword"

	err := ctx.BodyParser(&request)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	URL := "http://127.0.0.1:8020/auth/forgot-password"
	respCode, _ := util.HitMicroservicesAPI(FuncName, URL, "POST", "application/json", request)

	if respCode != 200 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", "Email sent succesfully")
}

func (c *Controller) CheckForgotPasswordToken(ctx *fiber.Ctx) error {
	var request dto.ForgotPasswordTokenRequest
	var getData entity.ResetPasswordToken
	FuncName := "CheckForgotPasswordToken"
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + ":" + msg)
		}

		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	count, err := c.Client.From("reset_password_tokens").Select("*", "", false).Eq("reset_password_token", request.Token).Single().ExecuteTo(&getData)

	if err != nil && count == 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	if getData.ExpiredAt.Before(time.Now()) {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusUnauthorized, cons.ErrTokenExpired, "")
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Token Active", nil)
}

func (c *Controller) ResetPassword(ctx *fiber.Ctx) error {
	var request dto.ResetPasswordRequest
	var getData entity.ResetPasswordToken
	FuncName := "ResetPassword"
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	count, err := c.Client.From("reset_password_tokens").Select("*", "", false).Eq("reset_password_token", request.Token).Single().ExecuteTo(&getData)

	if err != nil && count == 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusNotFound, cons.ErrTokenExpired, err.Error())
	}

	if getData.ExpiredAt.Before(time.Now()) {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusUnauthorized, cons.ErrTokenExpired, nil)
	}
	hashPassword, _ := util.HashPassword(request.Password)

	_, _, err = c.Client.From("users").Update(map[string]interface{}{
		"password":   string(hashPassword),
		"updated_at": time.Now(),
		"updated_by": getData.Email,
	}, "", "").Eq("email", getData.Email).Single().Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusBadGateway, cons.ErrFailed+" to update", err.Error())
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", "")
}

func (c *Controller) Logout(ctx *fiber.Ctx) error {
	FuncName := "Logout"
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		log.Error().Msg("API Endpoint /" + FuncName + " : Header Missing")
		return util.GenerateResponse(ctx, http.StatusBadRequest, "Authorization header missing", "")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		log.Error().Msg("API Endpoint /" + FuncName + ": Token Invalid Format")
		return util.GenerateResponse(ctx, http.StatusUnauthorized, "Invalid Token Format", "")
	}

	token := tokenParts[1]
	_, _, err := c.Client.From("users_token").Update(map[string]interface{}{
		"is_active": false,
	}, "", "").Eq("token", token).Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusInternalServerError, cons.ErrFailed+" to logout, Please try again", err.Error())
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Logout Success!", "")

}
