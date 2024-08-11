package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	launcher *launcher.Launcher
	page     *rod.Page
}

func Launch() *Browser {
	br := Browser{}

	br.launcher = launcher.New().Headless(true).Set("no-sandbox")
	u := br.launcher.MustLaunch()

	// Create blank page first:
	br.page = rod.New().ControlURL(u).MustConnect().MustPage()

	return &br
}

func (br *Browser) Close() {
	br.page.MustClose()
	br.launcher.Kill()
}

func (br *Browser) Navigate(url string) {
	br.page.MustNavigate(url)
	br.page.MustWaitIdle()
}

func (br *Browser) Click(elm string) {
	br.page.MustElement(elm).MustClick()
	br.page.MustWaitIdle()
}

func (br *Browser) Input(elm string, str string) {
	br.page.MustElement(elm).MustInput(str)
}

func (br *Browser) GetElmHTML(elm string) string {
	return br.page.MustElement(elm).MustHTML()
}

func (br *Browser) GetElmText(elm string) string {
	return br.page.MustElement(elm).MustText()
}

func (br *Browser) Screenshot(filename string) {
	br.page.MustScrollScreenshot(filename)
}
