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
	var key string
	var lang string

	flag.StringVar(&key, "key", "Greeting", "key to select the string")
	flag.StringVar(&lang, "lang", "en-US", "language to use")

	flag.Parse()

	// -

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load JSON message files directly
	if _, err := bundle.LoadMessageFile("translations/en_US.json"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading en bundle: %v\n", err)
		os.Exit(1)
	}
	if _, err := bundle.LoadMessageFile("translations/es_ES.json"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading es bundle: %v\n", err)
		os.Exit(1)
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	config := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: key,
	})

	fmt.Printf("%s\n", config)
}
