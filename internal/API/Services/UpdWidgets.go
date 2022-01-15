package Services

import (
	"github.com/fatih/color"
	"main/internal/Parsing"
)

func UpdWidgets() string {
	parsing := Parsing.Parsing()
	for _, end := range parsing {
		color.New(color.FgCyan).Add(color.Underline).Println(end)
	}
	return "Updated"
}
