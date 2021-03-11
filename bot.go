package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

var cd string

func main() {
	dt := time.Now()
	cd = dt.Format("02-01-2006")
	fmt.Println("Current date:", cd)
	//exchange()
	//wowapi()

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "",

		Token:  cfg.TelegramToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Приветствую тебя друг мой! Нужно больше /gold:)")
	})

	//buttons
	inlineBtn1 := tb.InlineButton{
		Unique: "wowtoken",
		Text:   "Курс WoW Token 🌞",
	}

	inlineBtn2 := tb.InlineButton{
		Unique: "bitcoin",
		Text:   "Курс Bitcoin 🌞",
	}

	inlineBtn3 := tb.InlineButton{
		Unique: "tenge",
		Text:   "Курс Тенге 🌞",
	}

	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Sender, "Wow token says 'Hi'!")
	})

	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Sender, "Биткоин says 'Hi'!")
	})

	b.Handle(&inlineBtn3, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Тенге says 'Hi'!")
	})

	inlineKeys := [][]tb.InlineButton{
		{inlineBtn1, inlineBtn2, inlineBtn3},
	}
	b.Handle("/gold", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Выберите нужное!",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})
	b.Start()

}
