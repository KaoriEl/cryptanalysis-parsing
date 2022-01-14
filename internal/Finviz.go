package internal

import (
	"github.com/fatih/color"
	"os"
)

func Finviz(urlGo string) {
	color.New(color.FgBlue).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("Finviz"))
	ctx, cancel := ChromeConfiguration()

	filenames := []string{
		"/Finviz" + RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div#canvas-wrapper",
	}

	Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
