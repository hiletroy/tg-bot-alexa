// -*- tab-width: 4; indent-tabs-mode: nil -*-

package main

import (
	"database/sql"
    "log"
    "os"

	_ "github.com/lib/pq"
    tb "gopkg.in/tucnak/telebot.v2"
)

// func dbFunc(db *sql.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
//             c.String(http.StatusInternalServerError,
//                 fmt.Sprintf("Error creating database table: %q", err))
//             return
//         }

//         if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
//             c.String(http.StatusInternalServerError,
//                 fmt.Sprintf("Error incrementing tick: %q", err))
//             return
//         }

//         rows, err := db.Query("SELECT tick FROM ticks")
//         if err != nil {
//             c.String(http.StatusInternalServerError,
//                 fmt.Sprintf("Error reading ticks: %q", err))
//             return
//         }

// 		c.String(http.StatusOK, fmt.Sprintf("<pre>ENV: %s\n</pre>", os.Environ()))
// 		//fmt.Fprintf(os.Stdout, "ENV: %s", os.Environ())

//         defer rows.Close()

//         for rows.Next() {
//             var tick time.Time
//             if err := rows.Scan(&tick); err != nil {
//                 c.String(http.StatusInternalServerError,
//                     fmt.Sprintf("Error scanning ticks: %q", err))
//                 return
//             }
//             c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
//         }
//     }
// }

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

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

    if _, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (id int primary key, chatid int, username text, messages text, date timestamp)");
    err != nil {
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
