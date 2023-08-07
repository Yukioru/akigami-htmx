package locales

import (
	"github.com/gofiber/fiber/v2"
)

func Header() fiber.Map {
	return fiber.Map{
		"menu": fiber.Map{
			"home": Localize("menu.home"),
		},
	}
}
