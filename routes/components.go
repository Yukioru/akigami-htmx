package routes

import (
	"akigami.co/controllers/components"
	"github.com/gofiber/fiber/v2"
)

func InitComponentsRoutes(router fiber.Router) {
	router.Get("/auth", components.AuthComponentController)
}
