package controller

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
	"github.com/mustafasegf/notion-note/util"
)

type Line struct {
	bot    *linebot.Client
	notion core.Notion
}

func NewLinkController(bot *linebot.Client, notion core.Notion) *Line {
	return &Line{
		bot:    bot,
		notion: notion,
	}
}

func (ctrl *Line) LineCallback(w http.ResponseWriter, req *http.Request) {
	events, err := ctrl.bot.ParseRequest(req)
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
				_, err = ctrl.notion.CreateNote(title, body)
				res := "created"
				if err != nil {
					log.Println(err)
					res = "failed"
				}
				if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
