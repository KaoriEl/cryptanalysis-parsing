package main

import (
	"github.com/fatih/color"
	"main/internal"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		internal.StartServer()
	}()
	parsing := internal.Parsing()
	for _, end := range parsing {
		color.New(color.FgCyan).Add(color.Underline).Println(end)
	}
	wg.Wait()
}
