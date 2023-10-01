package components

import (
	"time"

	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AuthComponentController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)

	email := c.FormValue("email")
	log.Info(email)

	time.Sleep(1 * time.Second)

	return utils.HTMLResponse(c, utils.HTMLResponseOptions{
		RouteType: "components",
		RouteKey:  "authFormSuccess",
		Locales: fiber.Map{
			"auth": locales.Auth(localizer),
		},
	})
}
