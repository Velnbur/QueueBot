package services

import (
	"strconv"

	"github.com/Velnbur/QueueBot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DaysView(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	var days [6]models.Day
	models.ListDays(upd.CallbackQuery.Data[1:], &days)

	var inlineKeyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(days); i++ {
		var inlineRow []tgbotapi.InlineKeyboardButton
		inlineRow = append(
			inlineRow,
			tgbotapi.NewInlineKeyboardButtonData(
				days[i].Data,
				"t"+strconv.Itoa(days[i].ID)),
		)
		inlineKeyboard = append(inlineKeyboard, inlineRow)
	}
	msg := tgbotapi.NewEditMessageText(
		upd.CallbackQuery.Message.Chat.ID,
		upd.CallbackQuery.Message.MessageID,
		"Days",
	)
	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}
	bot.Send(msg)
}

func TimesView(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	var queue []models.Queue
	models.ListQueue(upd.CallbackQuery.Data[1:], &queue)

	var inlineKeyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(queue); i++ {
		var inlineRow []tgbotapi.InlineKeyboardButton

		var msgText = queue[i].Time + "\xE2\x9C\x85"
		if queue[i].User.Name.Valid {
			msgText = queue[i].Time + "\xE2\x9D\x8C"
		}
		inlineRow = append(
			inlineRow,
			tgbotapi.NewInlineKeyboardButtonData(
				msgText,
				"p"+strconv.Itoa(queue[i].ID)),
		)
		inlineKeyboard = append(inlineKeyboard, inlineRow)
	}
	msg := tgbotapi.NewEditMessageText(
		upd.CallbackQuery.Message.Chat.ID,
		upd.CallbackQuery.Message.MessageID,
		"Time",
	)
	msg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard}
	bot.Send(msg)
}

func WeeksView(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	var weeks [5]models.Week
	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "Check weeks")

	models.ListWeeks(&weeks)

	var inlineKeyboard [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(weeks); i++ {
		var inlineRow []tgbotapi.InlineKeyboardButton
		inlineRow = append(
			inlineRow,
			tgbotapi.NewInlineKeyboardButtonData(
				weeks[i].Date,
				"d"+strconv.Itoa(weeks[i].ID)),
		)
		inlineKeyboard = append(inlineKeyboard, inlineRow)
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard...)
	bot.Send(msg)
}

func StartView(bot *tgbotapi.BotAPI, upd *tgbotapi.Update) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Show weeks"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ki"),
		),
	)
	models.AddUser(upd.Message.From.ID, upd.Message.From.UserName)

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}
