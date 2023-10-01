package utils

import (
	"strconv"
	"strings"
	"time"

	"akigami.co/locales"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type HTMLResponseOptions struct {
	Locales   fiber.Map
	Meta      MetadataInput
	RouteKey  string
	RouteType string
	LayoutKey string
	Data      fiber.Map
}

func HTMLResponse(c *fiber.Ctx, input HTMLResponseOptions) error {
	timer := time.Now()
	localizer := c.Locals("localizer").(*i18n.Localizer)
	localesMap := input.Locales
	if localesMap == nil {
		localesMap = fiber.Map{}
	}
	localesMap["header"] = locales.Header(localizer)
	localesMap["common"] = locales.Common(localizer)
	brandLocales := locales.Brand(localizer)

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
	metaInput.Brand.Title = brandLocales["title"].(string)
	metaInput.Brand.Description = brandLocales["description"].(string)

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

	c.Response().Header.Add("Server-Timing", "render_html;dur="+strconv.FormatInt(time.Since(timer).Milliseconds(), 10))

	return c.Render(input.RouteType+"/"+input.RouteKey, fiber.Map{
		"routeKey": input.RouteKey,
		"locales":  localesMap,
		"meta":     MakeMetadata(c, metaInput),
		"data":     input.Data,
		"ctx": fiber.Map{
			"hx": hx,
		},
	}, layoutKey)
}
