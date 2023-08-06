package utils

import "github.com/gofiber/fiber/v2"

func GetLayout(c *fiber.Ctx, name string) string {
	headers := c.GetReqHeaders()
	if headers["Hx-Request"] == "true" {
		return ""
	}

	return "layouts/" + name
}
