package pages

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func IndexPageController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)
	pageTitle := locales.Localize(localizer, "head.index")

	return utils.RenderHtml(c, utils.RenderHtmlInput{
		Meta: utils.MetadataInput{
			Title: pageTitle,
		},
	})
}
