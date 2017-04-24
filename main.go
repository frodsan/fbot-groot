package main

import (
	"log"
	"net/http"
	"os"

	"github.com/frodsan/fbot"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/webhook", fbot.Handler(createBot()))

	http.ListenAndServe(":"+port, nil)
}

func createBot() fbot.Bot {
	bot := fbot.NewBot()

	bot.VerifyToken = os.Getenv("VERIFY_TOKEN")

	return bot
}
