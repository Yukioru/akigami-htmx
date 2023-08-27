package utils

import (
	"github.com/gofiber/fiber/v2"
)

func GetLayout(c *fiber.Ctx, name string) string {
	hx := c.Get("Hx-Request")
	if hx == "true" {
		return ""
	}

	return "layouts/" + name
}
