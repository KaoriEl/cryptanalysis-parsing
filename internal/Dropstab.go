package internal

import (
	"github.com/fatih/color"
	"os"
)

func Dropstab(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("Dropstab"))
	ctx, cancel := ChromeConfiguration()
	filenames := []string{
		"/Dropstab-btcDominance" + RandStringRunes(10) + ".jpg",
		"/Dropstab-longsShorts" + RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div#btcDominance",
		"div#longsShorts",
	}
	Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
