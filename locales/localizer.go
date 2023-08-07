package locales

import (
	"github.com/bytedance/sonic"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Localizer *i18n.Localizer
var bundle *i18n.Bundle

func InitLocalizer() {
	bundle = i18n.NewBundle(language.Russian)

	bundle.RegisterUnmarshalFunc("json", sonic.Unmarshal)
	bundle.MustLoadMessageFile("locales/resources/ru.json")
	bundle.MustLoadMessageFile("locales/resources/en.json")

	Localizer = i18n.NewLocalizer(bundle, language.Russian.String(), language.English.String())
}

func Localize(messageID string) string {
	if Localizer == nil {
		InitLocalizer()
	}

	msg, err := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil {
		panic(err)
	}

	return msg
}
