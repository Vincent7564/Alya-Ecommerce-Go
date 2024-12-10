package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddProduct(ctx *fiber.Ctx) error {
	var request dto.AddProductRequest

	FuncName := "AddProduct :"
	err := ctx.BodyParser(&request)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
	}

	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		return cons.ErrValidationError
	}

	_, _, err = c.Client.From("products").Insert(map[string]interface{}{
		"product_name":        request.ProductName,
		"product_stock":       request.ProductStock,
		"product_price":       request.ProductPrice,
		"product_category_id": request.ProductCategoryId,
		"discount":            request.Discount,
		"description":         request.Description,
		"created_by":          "Testing",
	}, false, "", "", "").Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return cons.ErrSuccess
}
