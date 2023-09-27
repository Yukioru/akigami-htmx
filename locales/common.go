package locales

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Common(localizer *i18n.Localizer) fiber.Map {
	return fiber.Map{
		"loading": Localize(localizer, "common.loading"),
		"buttons": fiber.Map{
			"login": Localize(localizer, "common.buttons.login"),
		},
	}
}
