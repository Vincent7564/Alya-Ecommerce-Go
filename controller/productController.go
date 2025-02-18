package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddProduct(ctx *fiber.Ctx) error {
	var request dto.AddProductRequest

	FuncName := "AddProduct :"
	formData, err := ctx.MultipartForm()
	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	jsonStrArr, exists := formData.Value["data"]

	if !exists || len(jsonStrArr) == 0 {
		log.Error().Msg("Missing 'data' key in form-data")
		return cons.ErrInvalidRequest
	}

	if err = json.Unmarshal([]byte(jsonStrArr[0]), &request); err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	if errorMessage := util.ValidateData(&request); len(errorMessage) > 0 {
		for _, msg := range errorMessage {
			log.Error().Msg("Validation error in API Endpoint /" + FuncName + msg)
		}
		cons.ErrValidationError.Message += ": " + strings.Join(errorMessage, "; ")
		return cons.ErrValidationError
	}

	files := formData.File["product"] //Depends on key name
	var currentTime = time.Now()
	var productID int
	insertResult, _, err := c.Client.From("products").
		Insert(map[string]interface{}{
			"product_name":        request.ProductName,
			"product_stock":       request.ProductStock,
			"product_price":       request.ProductPrice,
			"product_category_id": request.ProductCategoryId,
			"discount":            request.Discount,
			"description":         request.Description,
			"created_by":          "Testing",
			"created_at":          currentTime.Format(time.RFC3339),
		}, false, "", "", "").Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	if len(insertResult) > 0 {
		var result []struct {
			ID int `json:"id"`
		}
		if err := json.Unmarshal(insertResult, &result); err == nil && len(result) > 0 {
			productID = result[0].ID
		}
	}

	if productID == 0 {
		log.Error().Msg("Failed to get inserted product ID")
		return cons.ErrInternalServerError
	}

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

		fileName := fmt.Sprintf("%s_%s_%s", request.ProductName, time.Now().Format("2006-01-02_15-04-05"), file.Filename)

		path, err := util.ImageUploader(fileBytes, fileName)

		if err != nil {
			return cons.ErrInternalServerError
		}
		src.Close()
		_, _, err = c.Client.From("product_images").Insert(map[string]interface{}{
			"product_id": productID,
			"images":     fileName,
			"created_at": time.Now(),
			"created_by": "Testing",
			"updated_at": time.Now(),
			"updated_by": "Testing",
		}, false, "", "", "").Execute()

		if err != nil {
			log.Error().Err(err).Msg("API Endpoint /" + FuncName)
			return cons.ErrInternalServerError
		}

		fmt.Println(path)
	}

	return cons.ErrSuccess
}

func (c *Controller) GetProduct(ctx *fiber.Ctx) error {
	FuncName := "GetProduct :"

	var products []entity.Products

	_, err := c.Client.From("products").
		Select("*,product_images(images) ,category(category_name)", "", false).
		ExecuteTo(&products)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	cdnURL := os.Getenv("NEXT_PUBLIC_PATH_SUPABASE")

	for i := range products {
		for j := range products[i].ProductImages {
			products[i].ProductImages[j].Images = cdnURL + products[i].ProductImages[j].Images
		}
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", products)
}

func (c *Controller) GetProductBySearch(ctx *fiber.Ctx) error {
	FuncName := "GetProductBySearch"
	var products []entity.Products

	productName := ctx.Query("product_name")
	categoryID := ctx.Query("category_id")

	query := c.Client.From("products").Select("*,category(category_name)", "", false)

	if productName != "" {
		query = query.Ilike("product_name", "%"+productName+"%")
	}
	if categoryID != "" {
		query = query.Eq("product_category_id", categoryID)
	}

	_, err := query.ExecuteTo(&products)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return util.GenerateResponse(ctx, http.StatusOK, "Success", products)
}

func (c *Controller) GetProductById(ctx *fiber.Ctx) error {
	FuncName := "GetProductById"
	var products entity.Products
	idParams := ctx.Params("id")

	_, err := c.Client.From("products").Select("*,category(category_name)", "", false).
		Eq("id", idParams).Single().ExecuteTo(&products)

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

func (c *Controller) DeleteProduct(ctx *fiber.Ctx) error {
	FuncName := "DeleteProduct"
	idParams := ctx.Params("id")

	_, _, err := c.Client.From("products").
		Delete("", "").
		Eq("id", idParams).
		Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}

func (c *Controller) UpdateProduct(ctx *fiber.Ctx) error {
	FuncName := "UpdateProduct"
	idParams := ctx.Params("id")

	var request dto.UpdateProductRequest

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

	_, _, err = c.Client.From("products").Update(map[string]interface{}{
		"product_name":        request.ProductName,
		"product_stock":       request.ProductStock,
		"product_price":       request.ProductPrice,
		"product_category_id": request.ProductCategoryId,
		"discount":            request.Discount,
		"description":         request.Description,
	}, "", "").Eq("id", idParams).Execute()

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

func (c *Controller) GetCategoryById(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	FuncName := "GetCategoryById"

	var categories entity.Categories

	_, err := c.Client.From("category").
		Select("id, category_name", "", false).
		Eq("is_active", "TRUE").
		Eq("id", idParams).
		Single().
		ExecuteTo(&categories)

	if err != nil {
		log.Error().Err(err).
			Str("function", FuncName).
			Msg("Failed to fetch categories from database")
		return cons.ErrDataNotFound
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
