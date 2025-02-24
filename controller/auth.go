package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func (c *Controller) InsertUser(ctx *fiber.Ctx) error {
	var user dto.InsertUserRequest
	var getData entity.UserEntity
	FuncName := "RegisterUser"
	err := ctx.BodyParser(&user)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	_, err = c.Client.From("users").Select("*", "", false).Eq("email", user.Email).Single().ExecuteTo(&getData)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrDataExisted
	}

	if getData.Email != "" {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrEmailExisted
	}

	_, err = c.Client.From("users").Select("*", "", false).Eq("username", user.Username).Single().ExecuteTo(&getData)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrDataExisted
	}

	if getData.Username != "" {
		log.Error().Msg("API Endpoint /" + FuncName)
		return cons.ErrUsernameExisted
	}

	if errorMessage := util.ValidateData(&user); len(errorMessage) > 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	hashPassword, _ := util.HashPassword(user.Password)

	_, _, err = c.Client.From("users").Insert(map[string]interface{}{
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
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusInternalServerError, cons.ErrFailed+" to insert user", err.Error())
	}

	return cons.ErrSuccess
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
	requestBody := map[string]string{
		"username": request.Username,
		"password": request.Password,
	}

	reqBody, _ := json.Marshal(requestBody)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(URL)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}
	err = client.Do(req, resp)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return util.GenerateResponse(ctx, http.StatusOK, "Success", response)
}

func (c *Controller) ForgotPassword(ctx *fiber.Ctx) error {
	var request dto.ForgotPasswordRequest
	var getData entity.UserEntity
	FuncName := "ForgotPassword"
	type data struct {
		Token       string
		ExpiredDate string
	}

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

	count, err := c.Client.From("users").Select("*", "", false).Eq("email", request.Email).Single().ExecuteTo(&getData)

	if err != nil && count == 0 {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrAccountNotFound
	}

	if token, err := util.GenerateRandomToken(); err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return util.GenerateResponse(ctx, http.StatusInternalServerError, cons.ErrFailed+" to generate token", err.Error())
	} else {
		d := data{}
		var tempTime = time.Now().Add(time.Minute * 30)
		d.Token = token
		d.ExpiredDate = tempTime.Format("03:04PM 2 January 2006")

		_, _, err = c.Client.From("reset_password_tokens").Insert(map[string]interface{}{
			"users_id":             getData.ID,
			"email":                getData.Email,
			"reset_password_token": d.Token,
			"created_at":           time.Now(),
			"expired_at":           d.ExpiredDate,
		}, false, "", "", "").Execute()

		if err != nil {
			log.Error().Err(err).Msg("API Endpoint /" + FuncName)
			return util.GenerateResponse(ctx, http.StatusBadGateway, cons.ErrFailed+" to insert token", err.Error())
		}

		tmpl, err := template.ParseFiles("utils/ForgotPassword.html")

		if err != nil {
			log.Error().Err(err).Msg("API Endpoint /" + FuncName)
			return util.GenerateResponse(ctx, http.StatusFailedDependency, cons.ErrFailed+" Parse Template", err.Error())
		}

		var body bytes.Buffer
		if err := tmpl.Execute(&body, d); err != nil {
			log.Error().Err(err).Msg("API Endpoint /" + FuncName)
			return util.GenerateResponse(ctx, http.StatusFailedDependency, cons.ErrFailed+" to execute template", err.Error())
		}

		util.SendEmail(getData.Email, "Forget password email for Alya Ecommerce Shop", body.String())

		return util.GenerateResponse(ctx, http.StatusOK, "Success", "")
	}
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
