package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/Velnbur/QueueBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
)

const DEBUG = true

func main() {
	var TelToken = flag.String("t", "", "Telegram Bot Token")
	flag.Parse()
	fmt.Println(TelToken)

	var err error

	models.DB, err = sql.Open("sqlite3", "queue.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer models.DB.Close()

	bot, err := tgbotapi.NewBotAPI(*TelToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				models.AddUser(update.Message.From.ID, update.Message.From.UserName)
			case "list":
				var weeks [5]models.Week
				models.ListWeeks(&weeks)
			}
		}
	}
}
