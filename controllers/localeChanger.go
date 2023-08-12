package controllers

import (
	"strings"

	"akigami.co/locales"
	"github.com/gofiber/fiber/v2"
)

func LocaleChangerController(c *fiber.Ctx) error {
	locale := c.Params("code")
	if !strings.Contains(strings.Join(locales.SupportedLanguages, " "), locale) {
		locale = locales.DefaultLanguage
	}

	c.Cookie(&fiber.Cookie{
		Name:     "locale",
		Value:    locale,
		MaxAge:   31536000, // 1 year
		HTTPOnly: true,
		Path:     "/",
	})

	referer := c.Get("Referer")

	redirectUrl := strings.Replace(referer, c.Protocol()+"://"+c.Hostname(), "", 1)
	if redirectUrl == "" {
		redirectUrl = "/"
	}

	return c.Redirect(redirectUrl)
}
