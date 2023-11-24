# Localizability (i18n + l10n) using goi18n

1. Import the following packages:

```
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
```
1. Install `goi18n` tool: `go install github.com/nicksnyder/go-i18n/v2/goi18n@latest`
2. Create an empty file to use for translation, for example `es_ES.json`, using `touch es_ES.json`
3. Run `go run main.go --lang=<'es' or 'en-US' without quotes> --key=<your key in bundle file>`

