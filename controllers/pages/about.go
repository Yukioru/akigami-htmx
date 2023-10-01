package pages

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AboutPageController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)
	pageTitle := locales.Localize(localizer, "head.about")

	return utils.HTMLResponse(c, utils.HTMLResponseOptions{
		Meta: utils.MetadataInput{
			Title: pageTitle,
			Breadcrumbs: utils.BreadcrumbsInput{
				{"/", locales.Localize(localizer, "head.index")},
				{c.Path(), pageTitle},
			},
		},
	})
}
