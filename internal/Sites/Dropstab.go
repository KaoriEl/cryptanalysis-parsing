package Sites

import (
	"github.com/fatih/color"
	"main/internal/Chrome"
	"main/internal/Extensions"
	"os"
)

func Dropstab(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("Dropstab"))
	ctx, cancel := Chrome.ChromeConfiguration(urlGo)
	filenames := []string{
		"/Dropstab-btcDominance" + Extensions.RandStringRunes(10) + ".jpg",
		"/Dropstab-longsShorts" + Extensions.RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div#btcDominance",
		"div#longsShorts",
	}
	Chrome.Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
