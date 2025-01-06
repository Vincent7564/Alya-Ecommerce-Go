package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddWishlist(ctx *fiber.Ctx) error {
	var request dto.AddWishlistRequest
	var tempData entity.Products
	FuncName := "AddWishlist :"
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}

	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	_, errors := c.Client.From("products").
		Select("*,category(category_name)", "", false).
		Eq("id", fmt.Sprintf("%v", request.ProductID)).
		Single().
		ExecuteTo(&tempData)

	if errors != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName + errors.Error())
		return cons.ErrDataNotFound
	}

	_, _, err = c.Client.From("wishlist").Insert(map[string]interface{}{
		"users_id":   request.UsersID,
		"product_id": request.ProductID,
		"created_at": time.Now(),
	}, false, "", "", "").Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return cons.ErrSuccess
}

func (c *Controller) GetWishlist(ctx *fiber.Ctx) error {
	FuncName := "GetWishlist :"
	var request dto.GetWishlistRequest
	var wishlist []entity.GetWishlist
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInvalidRequest
	}
	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	_, err = c.Client.From("wishlist").
		Select("id,users_id,product_id,created_at,products!inner(product_name,product_price,product_category_id)", "", false).
		Eq("users_id", fmt.Sprintf("%v", request.UsersID)).
		ExecuteTo(&wishlist)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", wishlist)
}

func (c *Controller) DeleteWishlist(ctx *fiber.Ctx) error {
	FuncName := "DeleteWishlist :"

	var idParams = ctx.Params("id")

	var tempData interface{}
	_, err := c.Client.From("wishlist").
		Select("*", "", false).
		Eq("id", fmt.Sprintf("%v", idParams)).
		Single().
		ExecuteTo(&tempData)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName + err.Error())
		return cons.ErrDataNotFound
	}

	_, _, err = c.Client.From("wishlist").Delete("", "").Eq("id", idParams).Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}
