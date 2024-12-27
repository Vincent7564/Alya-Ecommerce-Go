package router

import (
	"Alya-Ecommerce-Go/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

type Router struct {
	App    *fiber.App
	Client *supabase.Client
}

func (r *Router) Init() {
	controller := controller.Controller{Client: r.Client}

	PublicProductRoutes(r.App, controller)
	PublicEventRoutes(r.App, controller)

	AuthRoutes(r.App, controller)
	UserRoutes(r.App, controller)
	ProductRoutes(r.App, controller)
	EventRoutes(r.App, controller)
}

func InitRouter(app *fiber.App, client *supabase.Client) {
	router := Router{App: app, Client: client}
	router.Init()
}
