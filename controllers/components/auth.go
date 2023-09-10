package components

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AuthComponentController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)

	return utils.RenderHtml(c, utils.RenderHtmlInput{
		RouteType: "components",
		RouteKey:  "authFormSuccess",
		Locales: fiber.Map{
			"auth": locales.Auth(localizer),
		},
	})
}
