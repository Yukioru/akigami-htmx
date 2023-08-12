package utils

import (
	"github.com/gofiber/fiber/v2"
)

func GetLayout(c *fiber.Ctx, name string) string {
	boosted := c.Get("Hx-Boosted")
	if boosted == "true" {
		return ""
	}

	return "layouts/" + name
}
