package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func initializeBot() *tgbotapi.BotAPI {
	loadEnvError := godotenv.Load(".env")

	if loadEnvError != nil {
		log.Fatalf("Error loading .env file")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	return bot
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	bot.Send(msg)
}

func main() {
	bot := initializeBot()

	bot.Debug = true

	// create update bot configuration
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// create listener
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			userMsg := update.Message.Text
			userID := update.Message.Chat.ID

			log.Printf("[%v]: %s\n", userID, userMsg)

			switch userMsg {
			case "/start":
				sendMessage(bot, userID, "Welcome to the simple reply bot\ngithub repo: https://github.com/uaintknowme/simple-telegram-bot")

			case "author":
				sendMessage(bot, userID, "https://github.com/uaintknowme")

			case "hey":
				sendMessage(bot, userID, "hey hru?")
			}
		}

	}

}
