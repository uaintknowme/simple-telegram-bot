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

func main() {
	bot := initializeBot()

	bot.Debug = true
}
