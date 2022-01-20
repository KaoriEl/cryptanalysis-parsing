package Services

import (
	"main/internal/Structures"
	"main/pkg"
)

func GetWidgets() *Structures.WidgetScreenshots {
	files := pkg.OnlyReadDir("/var/www/investments-cryptanalysis-parsing/assets/img")
	var ctx = []string{
		"Dropstab-btcDominance",
		"Dropstab-longsShorts",
		"FearAndGreed",
		"FinanceYahoo",
		"Finviz",
		"FeIndex",
		"Cryptorank",
	}
	var screenshots *Structures.WidgetScreenshots
	m := map[string]string{}
	for _, f := range files.Files {
		for _, c := range ctx {
			m = FindSubstr(f, c, m)
		}

	}
	screenshots = &Structures.WidgetScreenshots{
		DropstabBtc:       m["Dropstab-btcDominance"],
		DropstabLongs:     m["Dropstab-longsShorts"],
		FearAndGreed:      m["FearAndGreed"],
		FinanceYahoo:      m["FinanceYahoo"],
		Finviz:            m["Finviz"],
		FearAndGreedIndex: m["FeIndex"],
		Cryptorank:        m["Cryptorank"],
	}
	return screenshots
}
