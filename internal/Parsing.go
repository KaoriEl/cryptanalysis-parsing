package internal

import (
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"sync"
)

type ParsingData struct {
	ImageUrl string
}

type Sites struct {
	Url []string
}

func initEnv() {
	filePrefix, _ := filepath.Abs("/var/www/investments-cryptanalysis-parsing/configs") // path from the working directory
	err := godotenv.Load(filePrefix + "/.env")
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "ENV was not loaded correctly"))
	}
}

func Parsing() []string {
	DeleteAll()
	initEnv()
	sites := Sites{
		Url: []string{
			os.Getenv("FearAndGreed"),
			os.Getenv("Finviz"),
			os.Getenv("FinanceYahoo"),
			os.Getenv("Dropstab"),
		},
	}

	var wg sync.WaitGroup
	for _, url := range sites.Url {
		urlGo := url
		wg.Add(1)
		switch urlGo {
		case os.Getenv("FearAndGreed"):
			go func() {
				defer wg.Done()
				FearAndGreed(urlGo)
			}()
		case os.Getenv("Finviz"):
			go func() {
				defer wg.Done()
				Finviz(urlGo)
			}()
		case os.Getenv("FinanceYahoo"):
			go func() {
				defer wg.Done()
				FinanceYahoo(urlGo)
			}()
		case os.Getenv("Dropstab"):
			go func() {
				defer wg.Done()
				Dropstab(urlGo)
			}()
		}

	}
	wg.Wait()

	return []string{
		"Finished parsing: " + os.Getenv("FearAndGreed"),
		"Finished parsing: " + os.Getenv("Finviz"),
		"Finished parsing: " + os.Getenv("FinanceYahoo"),
		"Finished parsing: " + os.Getenv("Dropstab"),
	}
}
