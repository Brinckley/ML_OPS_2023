package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"tgclient/pkg/telegram"
)

var (
	tgToken = "6747445825:AAGdhYbHr640Qv6pFpK3DGcMsMvu5pIw4Wo"
)

func main() {
	logger := logrus.Logger{Formatter: &logrus.JSONFormatter{}} // creating and initializing the logger
	logger.SetOutput(os.Stdout)

	bot, err := tgbotapi.NewBotAPI(tgToken) // creating connection to tg bot
	if err != nil {
		log.Fatalln("[ERR] Can't receive token")
	}
	bot.Debug = true

	tgBot := telegram.NewBot(bot, &logger) // creating custom bot client
	err = tgBot.Start()
	if err != nil {
		return
	}
}
