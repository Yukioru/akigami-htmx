package main

import (
	"log"
	"strings"

	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := jet.New("./views", ".jet")

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	app.Use(func(c *fiber.Ctx) error {
		cookieLang := c.Cookies("locale")
		if cookieLang == "" {
			cookieLang = locales.DefaultLanguage
			c.Cookie(&fiber.Cookie{
				Name:     "locale",
				Value:    cookieLang,
				MaxAge:   31536000, // 1 year
				HTTPOnly: true,
				Path:     "/",
			})
		}

		localizer := locales.InitLocalizer(cookieLang, c.Get("Accept-Language"))
		c.Locals("locale", cookieLang)
		c.Locals("localizer", localizer)
		return c.Next()
	})

	app.Static("/", "./public", fiber.Static{
		MaxAge: 31536000, // 1 year
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"locales": fiber.Map{
				"header": locales.Header(c.Locals("localizer").(*i18n.Localizer)),
			},
			"meta": utils.MakeMetadata(utils.MetadataInput{
				Locale:     c.Locals("locale").(string),
				Title:      "Главная",
				CurrentURL: "/",
			}),
		}, utils.GetLayout(c, "main"))
	})

	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.Render("pages/demo", fiber.Map{
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
	})

	app.Get("/about", func(c *fiber.Ctx) error {
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
	})

	app.Get("/locale/:code", func(c *fiber.Ctx) error {
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
	})

	app.Listen(":42069")
}
