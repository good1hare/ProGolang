package main

import (
	config "ProGolang/configs"
	"ProGolang/handler"
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	ConnectPostgres()

	bot, err := telegramApi.NewBotAPI(config.GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegramApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			handler.Respond(bot, update)
		}
	}
}
