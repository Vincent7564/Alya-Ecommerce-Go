package routes

import (
	"Alya-Ecommerce-Go/auth-service/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

type Router struct {
	App    *fiber.App
	Client *supabase.Client
}

func (r *Router) Init() {
	controller := controller.Controller{Client: r.Client}
	AuthRoutes(r.App, controller)
}

func InitRouter(app *fiber.App, client *supabase.Client) {
	router := Router{App: app, Client: client}
	router.Init()
}
