package controller

import (
	"Alya-Ecommerce-Go/model/dto"
	"Alya-Ecommerce-Go/model/entity"
	util "Alya-Ecommerce-Go/utils"
	cons "Alya-Ecommerce-Go/utils/const"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (c *Controller) AddEvent(ctx *fiber.Ctx) error {
	var request dto.AddEventRequest

	FuncName := "AddEvent :"
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

	_, _, err = c.Client.From("events").Insert(map[string]interface{}{
		"event_name":     request.EventName,
		"event_desc":     request.EventDescription,
		"event_start_at": request.EventStartAt,
		"event_end_at":   request.EventEndAt,
		"is_active":      request.IsActive,
		"created_at":     time.Now(),
		"created_by":     "API",
	}, false, "", "", "").Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}

	return cons.ErrSuccess
}

func (c *Controller) DeleteEvent(ctx *fiber.Ctx) error {
	idParams := ctx.Params("id")
	FuncName := "DeleteEvent"

	_, _, err := c.Client.From("events").Delete("", "").Eq("id", idParams).Execute()

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return cons.ErrSuccess
}

func (c *Controller) GetEvent(ctx *fiber.Ctx) error {
	FuncName := "GetEvent"

	var event []entity.Event

	_, err := c.Client.From("events").Select("*", "", false).ExecuteTo(&event)

	if err != nil {
		log.Error().Err(err).Msg("API Endpoint /" + FuncName)
		return cons.ErrInternalServerError
	}
	return util.GenerateResponse(ctx, http.StatusOK, "Success", event)
}
