// -*- tab-width: 4; -*-

package main

import (
    "log"
    "os"
    "fmt"

    tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
    var (
        port      = os.Getenv("PORT")
        token     = os.Getenv("TOKEN")
        publicURL = os.Getenv("HEROKU_APP_NAME") + ".herokuapp.com"
    )

    webhook := &tb.Webhook{
        Listen:   ":" + port,
        Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
    }

    pref := tb.Settings{
        Token:  token,
        Poller: webhook,
    }

    b, err := tb.NewBot(pref)

    if err != nil {
        log.Fatal(err)
    }

	inlineBtn1 := tb.InlineButton{
		Unique: "moon",
		Text:   "Moon 🌚",
	}

	inlineBtn2 := tb.InlineButton{
		Unique: "sun",
		Text:   "Sun 🌞",
	}

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn1, inlineBtn2},
	}

	b.Handle(&inlineBtn1, func(c *tb.Callback) {
        // Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
        // Send messages here
		b.Send(c.Sender, "Moon says 'Hi'!")
	})

	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Sun says 'Hi'!")
	})

	b.Handle("/hello", func(m *tb.Message) {
        b.Send(m.Sender, "You entered "+m.Text+" ("+m.Payload+")")
    })

	b.Handle("/pick_time", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Day or night, you choose",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

    b.Start()
}

// (setq indent-tabs-mode nil)
