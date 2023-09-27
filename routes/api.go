package routes

import (
	"akigami.co/controllers/api"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoutes(router fiber.Router) {
	router.Post("/auth", api.AuthApiController)
}
