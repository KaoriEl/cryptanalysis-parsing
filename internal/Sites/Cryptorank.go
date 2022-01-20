package Sites

import (
	"github.com/fatih/color"
	"main/internal/Chrome"
	"main/internal/Extensions"
	"os"
)

func Cryptorank(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("Cryptorank"))
	ctx, cancel := Chrome.ChromeConfiguration(urlGo)
	filenames := []string{
		"/Cryptorank" + Extensions.RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"canvas",
	}
	Chrome.Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
