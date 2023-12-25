package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	logger *logrus.Logger

	mainMenu       tgbotapi.InlineKeyboardMarkup
	additionalMenu tgbotapi.InlineKeyboardMarkup

	updates tgbotapi.UpdatesChannel
}

func NewBot(bot *tgbotapi.BotAPI, logger *logrus.Logger) *Bot {
	b := &Bot{
		bot:    bot,
		logger: logger,
	}
	err := b.initUpdatesChannel()
	if err != nil {
		return nil
	}
	return b
}

func (b *Bot) initUpdatesChannel() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	return b.setKeyboards()
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}
