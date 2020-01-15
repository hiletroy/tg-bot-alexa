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
        publicURL = os.Getenv("HEROKU_APP_NAME") + ".herokuapp.com" // you must add it to your config vars
        token     = os.Getenv("TOKEN")      // you must add it to your config vars
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

    fmt.Fprintf(os.Stderr, "ENV: %s", os.Environ())

    // b.Handle("/hello", func(m *tb.Message) {
	// 	b.Send(m.Sender, "Hi!")
    // })

    b.Handle("/hello", func(m *tb.Message) {
        b.Send(m.Sender, "You entered "+m.Text)
    })

    b.Start()
}

// (setq tab-width 4)
// (setq indent-tabs-mode nil)
