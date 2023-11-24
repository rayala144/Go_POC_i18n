package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:generate goi18n extract -sourceLanguage en

func addKeyToJSON(filePath, key, value string) error {
	// Read the exisiting JSON file
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %v", err)
	}

	// Unmarshal JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Add the new key-value pair to the map
	data[key] = value

	// Marshal the updated data back to JSON
	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Write the JSON data back to the file
	err = os.WriteFile(filePath, updatedJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing to JSON file: %v", err)
	}

	return nil
}

func main() {
	var key string
	var lang string
	var addKey bool

	enBundlePath := "translations/en_US.json"
	esBundlePath := "translations/es_ES.json"

	flag.StringVar(&key, "key", "Greeting", "key to select the string")
	flag.StringVar(&lang, "lang", "en-US", "language to use")
	flag.BoolVar(&addKey, "addKey", false, "add a key and value to budle")

	flag.Parse()

	// -
	if !addKey {
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

		// Load JSON message files directly
		if _, err := bundle.LoadMessageFile(enBundlePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading en bundle: %v\n", err)
			os.Exit(1)
		}
		if _, err := bundle.LoadMessageFile(esBundlePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading es bundle: %v\n", err)
			os.Exit(1)
		}

		localizer := i18n.NewLocalizer(bundle, lang)

		config := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: key,
		})

		fmt.Printf("%s\n", config)

	} else {

		fmt.Println("Enter the key: ")
		var userKeyAdd string
		fmt.Scanln(&userKeyAdd)

		fmt.Println("Enter value in en-US: ")
		var userValEn string
		scannerEn := bufio.NewScanner(os.Stdin)
		scannerEn.Scan()
		userValEn = scannerEn.Text()

		fmt.Println("Enter value in es: ")
		var userValEs string
		scannerEs := bufio.NewScanner(os.Stdin)
		scannerEs.Scan()
		userValEs = scannerEs.Text()

		if err := addKeyToJSON(enBundlePath, userKeyAdd, userValEn); err != nil {
			fmt.Printf("Error adding key to en-US bundle: %v\n", err)
		}

		if err := addKeyToJSON(esBundlePath, userKeyAdd, userValEs); err != nil {
			fmt.Printf("Error adding key to es bundle: %v\n", err)
		}

		fmt.Println("Key-value pairs added to the bundle files successfully.")

	}

}
