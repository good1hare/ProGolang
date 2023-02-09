package main

import (
	config "ProGolang/configs"
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	bot, err := telegramApi.NewBotAPI(config.GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegramApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := telegramApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
