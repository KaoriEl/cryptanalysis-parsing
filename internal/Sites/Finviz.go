package Sites

import (
	"github.com/fatih/color"
	"main/internal/Chrome"
	"main/internal/Extensions"
	"os"
)

func Finviz(urlGo string) {
	color.New(color.FgBlue).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("Finviz"))
	ctx, cancel := Chrome.ChromeConfiguration(urlGo)

	filenames := []string{
		"/Finviz" + Extensions.RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div#canvas-wrapper",
	}

	Chrome.Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
