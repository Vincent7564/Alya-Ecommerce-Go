package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(app *fiber.App, controller controller.Controller) {
	EventRoutes := app.Group("/event", middleware.ValidateTokenMiddleware(controller.Client))

	EventRoutes.Post("", controller.AddEvent)
	EventRoutes.Delete("/:id", controller.DeleteEvent)
	EventRoutes.Get("", controller.GetEvent)
}
