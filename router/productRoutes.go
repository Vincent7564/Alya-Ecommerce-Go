package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func PublicProductRoutes(app *fiber.App, controller controller.Controller) {
	PublicProductRoutes := app.Group("/product")

	PublicProductRoutes.Get("/category/", controller.GetCategory)
	PublicProductRoutes.Get("", controller.GetProduct)
	PublicProductRoutes.Get("/search", controller.GetProductBySearch)
	PublicProductRoutes.Get("/:id", controller.GetProductById)

}

func ProductRoutes(app *fiber.App, controller controller.Controller) {
	ProductRoutes := app.Group("/product", middleware.ValidateTokenMiddleware(controller.Client))

	ProductRoutes.Post("/category/", controller.AddCategory)
	// ProductRoutes.Get("/category/", controller.GetCategory)
	ProductRoutes.Get("/category/:id", controller.GetCategoryById)
	ProductRoutes.Patch("/category/:id", controller.UpdateCategory)
	ProductRoutes.Delete("/category/:id", controller.DeleteCategory)

	ProductRoutes.Post("", controller.AddProduct)
	// ProductRoutes.Get("", controller.GetProduct)
	ProductRoutes.Delete("/:id", controller.DeleteProduct)
	ProductRoutes.Patch("/:id", controller.UpdateProduct)
	// ProductRoutes.Get("/:id", controller.GetProductById)
}
