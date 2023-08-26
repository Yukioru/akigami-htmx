package pages

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func DemoPageController(c *fiber.Ctx) error {
	return c.Render("pages/demo", fiber.Map{
		"routeKey": "demo",
		"locales": fiber.Map{
			"header": locales.Header(c.Locals("localizer").(*i18n.Localizer)),
		},
		"meta": utils.MakeMetadata(utils.MetadataInput{
			Locale:     c.Locals("locale").(string),
			Title:      "Демо",
			CurrentURL: "/demo",
			Breadcrumbs: utils.BreadcrumbsInput{
				{"/", "Главная"},
				{"/demo", "Демо"},
			},
		}),
	}, utils.GetLayout(c, "main"))
}
