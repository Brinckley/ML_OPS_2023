package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) setKeyboards() error {
	b.setKeyboardMainMenu()
	b.setKeyboardYolovType()
	return nil
}

func (b *Bot) setKeyboardMainMenu() {
	b.mainMenu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Object Detection", "/objdetec"),
			tgbotapi.NewInlineKeyboardButtonData("Image Segmentation", "/imgsegm"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Pose Estimation", "/posestim"),
			tgbotapi.NewInlineKeyboardButtonData("Image Classification", "/imgclf"),
		),
	)
}

func (b *Bot) setKeyboardYolovType() {
	b.additionalMenu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yolov8n", "/yolov8n"),
			tgbotapi.NewInlineKeyboardButtonData("Yolov8s", "/yolov8s"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yolov8m", "/yolov8m"),
			tgbotapi.NewInlineKeyboardButtonData("Yolov8l", "/yolov8l"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Yolov8x", "/yolov8x"),
		),
	)
}
