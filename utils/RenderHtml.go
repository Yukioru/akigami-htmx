package utils

import (
	"strings"

	"akigami.co/locales"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type RenderHtmlInput struct {
	Locales   fiber.Map
	Meta      MetadataInput
	RouteKey  string
	RouteType string
	LayoutKey string
	Data      fiber.Map
}

func RenderHtml(c *fiber.Ctx, input RenderHtmlInput) error {
	localizer := c.Locals("localizer").(*i18n.Localizer)
	localesMap := input.Locales
	if localesMap == nil {
		localesMap = fiber.Map{}
	}
	localesMap["header"] = locales.Header(localizer)
	localesMap["common"] = locales.Common(localizer)

	currentPath := c.Path()
	if input.RouteKey == "" {
		input.RouteKey = strings.TrimPrefix(currentPath, "/")
	}
	if input.RouteKey == "" {
		input.RouteKey = "index"
	}

	if input.RouteType == "" {
		input.RouteType = "pages"
	}

	metaInput := MetadataInput{}
	metaInput.Locale = c.Locals("locale").(string)
	metaInput.CurrentURL = currentPath

	metaInput.Title = input.Meta.Title
	metaInput.Description = input.Meta.Description
	metaInput.Breadcrumbs = input.Meta.Breadcrumbs

	layoutKey := input.LayoutKey
	if layoutKey == "" {
		layoutKey = "main"
	}

	hx := c.Get("Hx-Request")
	if hx == "true" || input.RouteType == "components" {
		layoutKey = ""
	} else {
		layoutKey = "layouts/" + layoutKey
	}

	return c.Render(input.RouteType+"/"+input.RouteKey, fiber.Map{
		"routeKey": input.RouteKey,
		"locales":  localesMap,
		"meta":     MakeMetadata(metaInput),
		"data":     input.Data,
		"ctx": fiber.Map{
			"hx": hx,
		},
	}, layoutKey)
}
