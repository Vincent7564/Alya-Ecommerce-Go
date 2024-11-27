package router

import (
	"Alya-Ecommerce-Go/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, controller controller.Controller) {
	AuthRoutes := app.Group("/auth")
	AuthRoutes.Post("/register", controller.InsertUser)
}
