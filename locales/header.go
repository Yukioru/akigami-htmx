package locales

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Header(localizer *i18n.Localizer) fiber.Map {
	return fiber.Map{
		"menu": fiber.Map{
			"index": Localize(localizer, "menu.index"),
			"demo":  Localize(localizer, "menu.demo"),
			"about": Localize(localizer, "menu.about"),
		},
	}
}
