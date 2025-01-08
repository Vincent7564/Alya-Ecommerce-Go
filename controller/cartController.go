package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddCart(ctx *fiber.Ctx) error {
	var request dto.AddCartRequest
	FuncName := "AddCart :"

	if err := ctx.BodyParser(&request); err != nil {
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

	_, _, err := c.Client.From("carts").Insert(map[string]interface{}{
		"user_id":    request.UsersID,
		"product_id": request.ProductID,
		"qty":        request.Qty,
		"created_at": time.Now(),
		"created_by": nil,
		"updated_at": nil,
		"updated_by": nil,
	}, false, "", "", "").Execute()
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}

func (c *Controller) GetCart(ctx *fiber.Ctx) error {
	FuncName := "GetCart :"
	idParams := ctx.Params("id")
	var carts []dto.GetCartResponse

	_, err := c.Client.From("carts").
		Select("user_id,qty,products(product_name, product_price,product_stock)", "", false).
		Eq("user_id", idParams).
		ExecuteTo(&carts)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return util.GenerateResponse(ctx, http.StatusOK, "Success", carts)
}
