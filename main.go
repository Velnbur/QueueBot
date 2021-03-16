package main

import (
	"github.com/Velnbur/QueueBot/models"
	"database/sql"
	"flag"
	_ "github.com/mattn/go-sqlite3"
	tb "github.com/tucnak/telebot"
	"log"
	"time"
)

func ListDays(m *tb.Message, s *tb.ReplyMarkup) {

}

func main() {
	var TelToken = flag.String("t", "", "Telegram Bot Token")
	flag.Parse()

	var err error

	models.DB, err = sql.Open("sqlite3", "queue.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "", //"http://195.129.111.17:8012",

		Token:  *TelToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	var selector = &tb.ReplyMarkup{}
	var btn = selector.Data("\tU+21A2", "next")

	selector.Inline(selector.Row(btn))

	bot.Handle("/start", func(m *tb.Message) {
		_, err := bot.Send(m.Sender, "Okay, let's start")
		if err != nil {
			log.Fatal(err)
		}
	})

	bot.Handle("/list_days", func(m *tb.Message){
		_, err := bot.Send(m.Sender, "List for ...", selector)
		if err != nil {
			log.Fatal(err)
		}
	})

	bot.Start()
}
