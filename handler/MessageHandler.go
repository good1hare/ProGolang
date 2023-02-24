package handler

import (
	models "ProGolang/models"
	"ProGolang/storage"
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Respond(bot *telegramApi.BotAPI, update telegramApi.Update) {
	var user models.User
	userResult := storage.PostgreSQL().First(&user, "user_name = ?", update.Message.From.UserName)
	if userResult.Error != nil {
		storage.PostgreSQL().Create(&models.User{UserName: update.Message.From.UserName, ChatId: update.Message.Chat.ID})
	}
	message := telegramApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	message.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
