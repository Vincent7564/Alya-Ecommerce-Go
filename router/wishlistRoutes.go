package router

import (
	"Alya-Ecommerce-Go/controller"
	"Alya-Ecommerce-Go/middleware"

	"github.com/gofiber/fiber/v2"
)

func WishlistRoutes(app *fiber.App, controller controller.Controller) {
	WishlistRoutes := app.Group("/wishlist", middleware.ValidateTokenMiddleware(controller.Client))

	WishlistRoutes.Post("", controller.AddWishlist)
	WishlistRoutes.Get("/:id", controller.GetWishlist)
	WishlistRoutes.Delete("/:id", controller.DeleteWishlist)
}
