package api

import (
	"time"

	"akigami.co/locales"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func AuthApiController(c *fiber.Ctx) error {
	hx := c.Get("Hx-Request")

	time.Sleep(1 * time.Second)

	if hx == "true" {
		return c.Redirect("/components/auth", 303)
	} else {
		localizer := c.Locals("localizer").(*i18n.Localizer)
		authLocales := locales.Auth(localizer)

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": authLocales["confirm"].(fiber.Map)["text"],
		})
	}
}
