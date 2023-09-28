package locales

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Brand(localizer *i18n.Localizer) fiber.Map {
	return fiber.Map{
		"title":       Localize(localizer, "brand.title"),
		"description": Localize(localizer, "brand.text"),
	}
}
