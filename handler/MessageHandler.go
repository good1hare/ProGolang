package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Respond(bot *telegramApi.BotAPI, update telegramApi.Update) {
	message := telegramApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	message.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
