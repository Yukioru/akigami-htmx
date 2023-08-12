package locales

import (
	"github.com/bytedance/sonic"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

var DefaultLanguage = language.Russian.String()
var SupportedLanguages = []string{
	language.Russian.String(),
	language.English.String(),
}

func InitLocalizer(lang string, accept string) *i18n.Localizer {
	if bundle == nil {
		bundle = i18n.NewBundle(language.Russian)

		bundle.RegisterUnmarshalFunc("json", sonic.Unmarshal)
		bundle.MustLoadMessageFile("locales/resources/ru.json")
		bundle.MustLoadMessageFile("locales/resources/en.json")
	}

	return i18n.NewLocalizer(bundle, lang, accept)
}

func Localize(localizer *i18n.Localizer, messageID string) string {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil {
		panic(err)
	}

	return msg
}
