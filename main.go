package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:generate goi18n extract -sourceLanguage en

func main() {
	var count int
	var lang string

	flag.IntVar(&count, "count", 0, "number of items to buy")
	flag.StringVar(&lang, "lang", "en", "language to use")

	flag.Parse()

	// -

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load JSON message files directly
	if _, err := bundle.LoadMessageFile("translations/en_US.json"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading en bundle: %v\n", err)
		os.Exit(1)
	}
	if _, err := bundle.LoadMessageFile("translations/es_SP.json"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading es bundle: %v\n", err)
		os.Exit(1)
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	// buying := localizer.MustLocalize(&i18n.LocalizeConfig{
	// 	DefaultMessage: &i18n.Message{
	// 		ID:    "DefaultMessage",
	// 		One:   "You're buying 1 cookie.",
	// 		Other: "You're buying {{.PluralCount}} cookies.",
	// 	},
	// 	PluralCount: count,
	// })

	greeting := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Greeting",
		},
		PluralCount: count,
	})

	// fmt.Printf("%s\n", buying)
	fmt.Printf("%s\n", greeting)
}
