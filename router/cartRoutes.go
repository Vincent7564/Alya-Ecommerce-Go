package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(app *fiber.App, controller controller.Controller) {
	CartRoutes := app.Group("/cart", middleware.ValidateTokenMiddleware(controller.Client))

	CartRoutes.Post("", controller.AddCart)
	CartRoutes.Get("/:id", controller.GetCart)
	CartRoutes.Delete("/:id", controller.DeleteCartItem)
}
