package main

import (
	"log"
	"net/http"

	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
	"github.com/mustafasegf/notion-note/util"
)

func main() {
	config, _ := util.LoadConfig()
	client := notionapi.NewClient(notionapi.Token(config.NotionToken))

	notion := core.Notion{
		Config: config,
		Client: client,
	}

	bot, err := linebot.New(config.LineSecret, config.LineToken)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					title, body := util.ParseText(message.Text)
					_, err = notion.CreateNote(title, body)
					res := "created"
					if err != nil {
						log.Println(err)
						res = "failed"
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
