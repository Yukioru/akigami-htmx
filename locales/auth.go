package locales

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Auth(localizer *i18n.Localizer) fiber.Map {
	return fiber.Map{
		"title":  Localize(localizer, "auth.title"),
		"submit": Localize(localizer, "auth.submit"),
		"confirm": fiber.Map{
			"title": Localize(localizer, "auth.confirm.title"),
			"text":  Localize(localizer, "auth.confirm.text"),
		},
		"email": fiber.Map{
			"title":       Localize(localizer, "auth.email.title"),
			"placeholder": Localize(localizer, "auth.email.placeholder"),
		},
	}
}
