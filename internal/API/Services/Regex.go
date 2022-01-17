package Services

import (
	"regexp"
)

func FindSubstr(files string, find string, m map[string]string) map[string]string {
	var re = regexp.MustCompile(`\/assets\/img\/(` + find + `)`)
	var str = files

	if len(re.FindStringIndex(str)) > 0 {
		switch re.FindStringSubmatch(str)[1] {
		case "Dropstab-btcDominance":
			m["Dropstab-btcDominance"] = files
			break
		case "Dropstab-longsShorts":
			m["Dropstab-longsShorts"] = files
			break
		case "FearAndGreed":
			m["FearAndGreed"] = files
			break
		case "FinanceYahoo":
			m["FinanceYahoo"] = files
			break
		case "Finviz":
			m["Finviz"] = files
			break
		}
	}
	return m
}
