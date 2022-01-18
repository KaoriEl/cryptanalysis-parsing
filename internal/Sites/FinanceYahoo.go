package Sites

import (
	"github.com/fatih/color"
	"main/internal/Chrome"
	"main/internal/Extensions"
	"os"
)

func FinanceYahoo(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("FinanceYahoo"))
	ctx, cancel := Chrome.ChromeConfiguration(urlGo)

	filenames := []string{
		"/FinanceYahoo" + Extensions.RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div#Lead-3-FinanceHeader-Proxy > div",
	}

	Chrome.Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
