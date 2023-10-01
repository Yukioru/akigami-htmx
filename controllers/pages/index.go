package pages

import (
	"akigami.co/db/models"
	"akigami.co/locales"
	"akigami.co/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func IndexPageController(c *fiber.Ctx) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)
	pageTitle := locales.Localize(localizer, "head.index")

	userModel := models.Get[models.UserSchema](c, "users")

	users := userModel.Find()
	log.Info(users)

	return utils.HTMLResponse(c, utils.HTMLResponseOptions{
		Meta: utils.MetadataInput{
			Title: pageTitle,
		},
	})
}
