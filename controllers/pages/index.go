package pages

import (
	"context"

	"akigami.co/db"
	"akigami.co/db/models"
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mongodb.org/mongo-driver/bson"
)

func IndexPageController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)
	pageTitle := locales.Localize(localizer, "head.index")

	cursor, err := db.DB.Collection("users").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return utils.RenderHtml(c, utils.RenderHtmlInput{
		Meta: utils.MetadataInput{
			Title: pageTitle,
		},
	})
}
