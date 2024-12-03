package router

import (
	"Alya-Ecommerce-Go/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, controller controller.Controller) {
	AuthRoutes := app.Group("/auth")

	AuthRoutes.Post("/register", controller.InsertUser)
	AuthRoutes.Post("/login", controller.Login)
	AuthRoutes.Post("/forgot-password", controller.ForgotPassword)
	AuthRoutes.Post("/check-password-token", controller.CheckForgotPasswordToken)
	AuthRoutes.Post("/reset-password", controller.ResetPassword)
	AuthRoutes.Post("/logout", controller.Logout)
}
