package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddProduct(ctx *fiber.Ctx) error {
	var request dto.AddProductRequest

	FuncName := "AddProduct :"
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

func (c *Controller) GetProduct(ctx *fiber.Ctx) error {
	FuncName := "GetProduct :"

	var products []entity.Products

	_, err := c.Client.From("products").
		Select("*, category(category_name)", "", false).
		ExecuteTo(&products)
	fmt.Print(products)
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", products)
}

func (c *Controller) AddCategory(ctx *fiber.Ctx) error {
	var request dto.AddCategoryRequest
	FuncName := "AddCategory :"

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

	_, _, err = c.Client.From("category").Insert(map[string]interface{}{
		"category_name": request.CategoryName,
		"is_active":     true,
	}, false, "", "", "").Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}

func (c *Controller) GetCategory(ctx *fiber.Ctx) error {
	FuncName := "GetCategory"

	var categories []entity.Categories

	_, err := c.Client.From("category").
		Select("id, category_name", "", false).
		Eq("is_active", "TRUE").
		ExecuteTo(&categories)

	if err != nil {
		log.Error().Err(err).
			Str("function", FuncName).
			Msg("Failed to fetch categories from database")
		return cons.ErrInternalServerError
	}

	if len(categories) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No categories found",
		})
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Categories retrieved successfully", categories)
}

func (c *Controller) UpdateCategory(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	FuncName := "UpdateCategory"
	var request dto.UpdateCategoryRequest

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
	_, _, err = c.Client.From("category").Update(map[string]interface{}{
		"category_name": request.CategoryName,
	}, "", "").Eq("id", string(idParam)).Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint / " + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Update Category Success", nil)
}

func (c *Controller) DeleteCategory(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	FuncName := "DeleteCategory"
	_, _, err := c.Client.From("category").Update(map[string]interface{}{
		"is_active": false,
	}, "", "").Eq("id", idParam).Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint / " + FuncName)
		return cons.ErrInternalServerError
	}
	return util.GenerateResponse(ctx, http.StatusOK, "Success", nil)
}
