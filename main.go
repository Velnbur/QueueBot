package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/Velnbur/QueueBot/models"
	srvc "github.com/Velnbur/QueueBot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var TelToken = flag.String("t", "", "Telegram Bot Token")
	flag.Parse()
	fmt.Println(TelToken)

	var err error

	models.DB, err = sql.Open("sqlite3", "Bot.sqlite")
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
		if update.Message == nil && update.CallbackQuery == nil { // ignore any non-Message Updates
			continue
		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data[0] {
			case 'd':
				srvc.DaysView(bot, &update)

			case 't':
				srvc.TimesView(bot, &update)

			}
		}
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				srvc.StartView(bot, &update)
			case "list_weeks":
				srvc.WeeksView(bot, &update)
			}
		}
	}
}
