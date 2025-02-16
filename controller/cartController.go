package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"fmt"
	"io"
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

func (c *Controller) DeleteCartItem(ctx *fiber.Ctx) error {
	FuncName := "DeleteCartItem :"
	idParams := ctx.Params("id")
	var cart entity.Cart
	_, err := c.Client.From("carts").Select("*", "", false).Eq("id", idParams).Single().ExecuteTo(&cart)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrDataNotFound
	}

	_, _, err = c.Client.From("carts").Delete("", "").Eq("id", idParams).Execute()
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return cons.ErrSuccess
}

func (c *Controller) UpdateCart(ctx *fiber.Ctx) error {
	FuncName := "UpdateCart :"
	idParams := ctx.Params("id")
	var request dto.UpdateCartRequest
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

	_, _, err := c.Client.From("carts").Update(map[string]interface{}{
		"qty":        request.Qty,
		"updated_at": time.Now(),
		"updated_by": nil,
	}, "", "").Eq("id", idParams).Execute()
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}

func (c *Controller) UploadImageTest(ctx *fiber.Ctx) error {
	FuncName := "UploadImageTest :"

	formData, err := ctx.MultipartForm()
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	files := formData.File["product"] //Depends on key name

	for _, file := range files {
		fmt.Println(file.Filename)

		src, err := file.Open()
		if err != nil {
			log.Error().Err(err).Msg("API Endpoint /" + FuncName)
			return cons.ErrInternalServerError
		}

		fileBytes, err := io.ReadAll(src)
		if err != nil {
			return cons.ErrInternalServerError
		}

		fileName := fmt.Sprintf("%s_%s_%s", "test", time.Now().Format("2006-01-02_15-04-05"), file.Filename)

		path, err := util.ImageUploader(fileBytes, fileName)

		if err != nil {
			return cons.ErrInternalServerError
		}
		src.Close()

		fmt.Println(path)
	}
	return cons.ErrSuccess

}
