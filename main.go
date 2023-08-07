package main

import (
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(compress.New())
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))

	app.Static("/", "./public", fiber.Static{
		Compress: true,
		MaxAge:   31536000, // 1 year
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"locales": fiber.Map{
				"header": locales.Header(),
			},
			"meta": utils.MakeMetadata(utils.Metadata{
				Title: "Главная",
			}),
		}, utils.GetLayout(c, "main"))
	})

	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.Render("pages/demo", fiber.Map{
			"meta": utils.MakeMetadata(utils.Metadata{
				Title: "Демо",
			}),
		}, utils.GetLayout(c, "main"))
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{
			"meta": utils.MakeMetadata(utils.Metadata{
				Title: "О нас",
			}),
		}, utils.GetLayout(c, "main"))
	})

	app.Listen(":42069")
}
