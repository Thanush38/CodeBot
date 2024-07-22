package main

import (
	bot "CodeBot/Bot"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run() // call the run function of bot/bot.go

	//Requests.Run()

}
