package Chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"main/internal/Extensions"
	"os"
)

func ChromeConfiguration(url string) (context.Context, context.CancelFunc) {
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome configuration options...")
	go Extensions.Progress(5)
	var imagesEnabled string
	if url == os.Getenv("FinanceYahoo") {
		imagesEnabled = "imagesEnabled=false"
	} else {
		imagesEnabled = "imagesEnabled=true"
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("blink-settings", imagesEnabled),
	)
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome NewExecAllocator...")

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	color.New(color.FgHiWhite).Add(color.Bold).Println("Сhrome context generate...")
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	return ctx, cancel
}
