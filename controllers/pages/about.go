package pages

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AboutPageController(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{
		"locales": fiber.Map{
			"header": locales.Header(c.Locals("localizer").(*i18n.Localizer)),
		},
		"meta": utils.MakeMetadata(utils.MetadataInput{
			Locale:     c.Locals("locale").(string),
			Title:      "О нас",
			CurrentURL: "/about",
			Breadcrumbs: utils.BreadcrumbsInput{
				{"/", "Главная"},
				{"/about", "О нас"},
			},
		}),
	}, utils.GetLayout(c, "main"))
}
