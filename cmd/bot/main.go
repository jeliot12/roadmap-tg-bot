package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jeliot12/psychologyBot/pkg/telegram"
)

func main() {
	TOKEN := "6252370167:AAHDGNMGezg6ozzOsDrWJLE5PkEEhgdkbBc"
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
