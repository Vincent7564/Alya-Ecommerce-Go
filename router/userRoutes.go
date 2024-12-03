package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, controller controller.Controller) {
	UserRoutes := app.Group("/users", middleware.ValidateTokenMiddleware(controller.Client))

	UserRoutes.Post("/check-password", controller.CheckPassword)
}
