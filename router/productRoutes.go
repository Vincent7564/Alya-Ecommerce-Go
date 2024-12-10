package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, controller controller.Controller) {
	ProductRoutes := app.Group("/product", middleware.ValidateTokenMiddleware(controller.Client))

	ProductRoutes.Post("", controller.AddProduct)
}
