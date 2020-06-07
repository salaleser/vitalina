package main

import (
	"github.com/salaleser/vitalina/bot"
	"github.com/salaleser/vitalina/util"
)

func main() {
	util.ReadConfig()

	util.InitLangaugeDetection()
	// util.InitScraper()

	bot.Start("Bot " + util.Config["token"])
}
