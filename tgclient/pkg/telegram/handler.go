package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"tgclient/pkg/sender"
)

const (
	startCmd = "start"
	objdetec = "objdetec"
	imgsegm  = "imgsegm"
	posestim = "posestim"
	imgclf   = "imgclf"
	yolov8n  = "yolov8n"
	yolov8s  = "yolov8s"
	yolov8m  = "yolov8m"
	yolov8l  = "yolov8l"
	yolov8x  = "yolov8x"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	sender.SendPostURL(message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, "Dataset url got. Wait for the answer....")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCmd:
		return b.handleStartCommand(message)
	case objdetec:
		return b.handleOperationType(message)
	case imgsegm:
		return b.handleOperationType(message)
	case posestim:
		return b.handleOperationType(message)
	case imgclf:
		return b.handleOperationType(message)
	case yolov8n:
		return b.handleYolov8Type(message)
	case yolov8s:
		return b.handleYolov8Type(message)
	case yolov8m:
		return b.handleYolov8Type(message)
	case yolov8l:
		return b.handleYolov8Type(message)
	case yolov8x:
		return b.handleYolov8Type(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hi")
	msg.ReplyMarkup = b.mainMenu
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleOperationType(message *tgbotapi.Message) error {
	sender.SendPostOperation(message.Command())
	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose your Yolov8 : ")
	msg.ReplyMarkup = b.additionalMenu
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleYolov8Type(message *tgbotapi.Message) error {
	sender.SendPostType(message.Command())
	msg := tgbotapi.NewMessage(message.Chat.ID, "Wait for the result")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")
	_, err := b.bot.Send(msg)
	return err
}
