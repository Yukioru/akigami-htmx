package main

import (
	"log"
	"os"

	"akigami.co/locales"
	"akigami.co/routes"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	isProduction := os.Getenv("GO_ENV") == "production"
	engine := jet.New("./views", ".jet")

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Prefork:     isProduction,
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

	routes.InitRoutes(app)

	app.Listen("127.0.0.1:42069")
}
