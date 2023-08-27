package routes

import (
	"akigami.co/controllers"
	"akigami.co/controllers/pages"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(router fiber.Router) {
	router.Get("/", pages.IndexPageController)
	router.Get("/auth", pages.AuthPageController)
	router.Get("/demo", pages.DemoPageController)
	router.Get("/about", pages.AboutPageController)

	router.Get("/locale/:code", controllers.LocaleChangerController)

	components := router.Group("/components")
	InitComponentsRoutes(components)
}
