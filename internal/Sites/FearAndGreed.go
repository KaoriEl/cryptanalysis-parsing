package Sites

import (
	"github.com/fatih/color"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/pkg/errors"
	"main/internal/Extensions"
	"main/internal/Structures"
	"os"
	"regexp"
)

func FearAndGreed(urlGo string) {
	color.New(color.FgHiMagenta).Add(color.Bold).Println("I start scraping the site: " + os.Getenv("FearAndGreed"))
	geziyor.NewGeziyor(&geziyor.Options{
		UserAgent:         "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		RobotsTxtDisabled: true,
		StartURLs:         []string{urlGo},
		ParseFunc:         ParsingFearAndGreed,
	}).Start()
}

func ParsingFearAndGreed(g *geziyor.Geziyor, r *client.Response) {
	if href, ok := r.HTMLDoc.Find("div#needleChart").Attr("style"); ok {
		var re = regexp.MustCompile(`(?:\(['"]?)(.*?)(?:['"]?\))`)
		var str = href
		if len(re.FindStringIndex(str)) > 0 {
			data := &Structures.ParsingData{ImageUrl: re.FindStringSubmatch(str)[1]}
			color.New(color.FgGreen).Add(color.Bold).Println("Image url: " + data.ImageUrl)
			color.New(color.FgHiMagenta).Add(color.Bold).Println("Starting file upload...")
			fileName := "/FearAndGreed" + Extensions.RandStringRunes(10) + ".jpg"
			Extensions.Index(data.ImageUrl, fileName)
		} else {
			color.New(color.FgRed).Add(color.Underline).Println(errors.New("Couldn't find image address"))
		}
	} else {
		color.New(color.FgRed).Add(color.Underline).Println(errors.New("Couldn't find the element from which to parse the link"))
	}
}
