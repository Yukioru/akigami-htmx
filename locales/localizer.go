package locales

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/bytedance/sonic"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func Localizer()  {
	bundle = i18n.NewBundle(language.Russian)

	bundle.RegisterUnmarshalFunc("json", sonic.Unmarshal)
	bundle.MustLoadMessageFile("locales/resouces/ru.json")
	bundle.MustLoadMessageFile("locales/resouces/en.json")

	localizer = i18n.NewLocalizer(bundle, language.Russian.String(), language.English.String())
}