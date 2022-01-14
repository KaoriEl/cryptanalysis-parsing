package internal

import (
	"github.com/fatih/color"
	"os"
)

func FinanceYahoo(urlGo string) {
	color.New(color.FgHiYellow).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("FinanceYahoo"))
	ctx, cancel := ChromeConfiguration()

	filenames := []string{
		"/FinanceYahoo" + RandStringRunes(10) + ".jpg",
	}
	elements := []string{
		"ul.Carousel-Slider",
	}

	Screeenshot(ctx, urlGo, filenames, elements, cancel)
}
