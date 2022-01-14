package internal

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
	"time"
)

func Screeenshot(ctx context.Context, urlGo string, filenames []string, elements []string, cancel context.CancelFunc) {
	go Progress(10)
	color.New(color.FgHiWhite).Add(color.Bold).Println("Defining filePrefix...")
	filePrefix, _ := filepath.Abs("/var/www/investments-cryptanalysis-parsing/assets/img") // path from the working directory
	color.New(color.FgHiWhite).Add(color.Bold).Println("Create buffer...")
	var imageBuf []byte
	color.New(color.FgHiWhite).Add(color.Bold).Println("Run chrome with context...")
	for key, element := range elements {
		if err := chromedp.Run(ctx, elementScreenshot(urlGo, element, &imageBuf)); err != nil {
			color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
		} else {
			color.New(color.FgHiWhite).Add(color.Bold).Println("Complete, close chrome...")
		}

		if err := ioutil.WriteFile(filePrefix+filenames[key], imageBuf, 0644); err != nil {
			color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't save screenshot"))
		} else {
			color.New(color.FgHiWhite).Add(color.Bold).Println("Complete, save screenshot. Exit func...")
		}
	}
	defer cancel()

}

func elementScreenshot(url string, sel string, imageBuf *[]byte) chromedp.Tasks {
	color.New(color.FgHiWhite).Add(color.Bold).Println("Taking a screenshot of an area...")
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second * 5),
		chromedp.Screenshot(sel, imageBuf, chromedp.NodeVisible),
	}
}

// AllPageSreenshot
//Функционало на будущее, пока что нах не нужно, можно ток сайты потестить заходити или нет ваще.
///**
func AllPageSreenshot(url string, imageBuf *[]byte) chromedp.Tasks {
	color.New(color.FgHiWhite).Add(color.Bold).Println("Taking a screenshot of all page...")
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(imageBuf, 100),
	}
}
