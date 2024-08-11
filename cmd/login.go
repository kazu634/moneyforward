package cmd

import (
	"log"

	"gitea.kazu634.com/kazu634/moneyforward/internal/lib/browser"
	"github.com/spf13/viper"
)

func init() {
	filter := logInit()
	log.SetOutput(filter)
}

func login(browser *browser.Browser) *browser.Browser {
	logInfo("Navigating to https://moneyforward.com/me")
	browser.Navigate("https://moneyforward.com/me")

	logInfo("Navigating to sign-in page")
	browser.Click("a[href^='/sign_in']")

	logInfo("Input Email address")
	browser.Input("input[type=email]", viper.Get("user").(string))
	browser.Click("#submitto")

	logInfo("Input password")
	browser.Input("input[name='mfid_user[password]']", viper.Get("password").(string))
	browser.Click("#submitto")

	return browser
}
