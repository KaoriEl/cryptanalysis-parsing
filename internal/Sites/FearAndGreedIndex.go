package Sites

import (
	"github.com/fatih/color"
	"main/internal/Chrome"
	"main/internal/Extensions"
	"os"
)

func FearAndGreedIndex(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("FearAndGreedIndex"))
	ctx, cancel := Chrome.ChromeConfiguration(urlGo)
	filenames := []string{
		"/FearAndGreedIndex" + Extensions.RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"div.columns > div:nth-child(2)",
	}
	Chrome.Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
