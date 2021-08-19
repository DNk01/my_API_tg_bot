package main

import (
	"github.com/DNk01/my_bot/pkg"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Your awesome TG token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	/*
		Place for DB
	*/

	telegramBot:=pkg.NewBot(bot)
	telegramBot.Start()
	if err!=nil{
		log.Fatal(err)
	}
}

