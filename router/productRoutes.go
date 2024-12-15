package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, controller controller.Controller) {
	ProductRoutes := app.Group("/product", middleware.ValidateTokenMiddleware(controller.Client))

	ProductRoutes.Post("", controller.AddProduct)
	ProductRoutes.Post("/category/", controller.AddCategory)
	ProductRoutes.Get("/category/", controller.GetCategory)
	ProductRoutes.Patch("/category/:id", controller.UpdateCategory)
	ProductRoutes.Delete("/category/:id", controller.DeleteCategory)
}
