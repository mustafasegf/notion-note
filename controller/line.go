package controller

import (
	"log"
	"net/http"

	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/service"
	"github.com/mustafasegf/notion-note/util"
)

type Line struct {
	svc *service.Line
	bot *linebot.Client
}

func NewLineController(bot *linebot.Client, svc *service.Line) *Line {
	return &Line{
		svc: svc,
		bot: bot,
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
				if !util.CheckIfCommand(message.Text) {
					return
				}
				userID := event.Source.UserID
				switch util.GetCommand(message.Text) {
				case "note":
					token := util.Tokenizer(message.Text)
					title, body := token["note"], token["body"]
					tags, ok := token["tags"]
					if !ok {
						tags = ""
					}
					res, err := ctrl.svc.CreateNote(userID, title, body, tags)
					if err != nil {
						log.Println(err)
					}
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "add":
					token := util.Tokenizer(message.Text)
					add, ok := token["add"]
					if !ok {
						add = ""
					}
					body := token["body"]

					var page *notionapi.DatabaseQueryResponse
					var pageID notionapi.ObjectID
					res := "failed"
					if add == "" {
						page, err = ctrl.svc.GetLatestNote(userID)
					} else {
						page, err = ctrl.svc.FindNote(userID, add)
					}

					if err == nil {
						pageID = page.Results[0].ID
						res, err = ctrl.svc.AppendNote(userID, string(pageID), body)
					}

					if err != nil {
						res = err.Error()
						log.Println(err)
					}
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "token":
					token := util.ParseTextOne(message.Text)
					res := ctrl.svc.UpdateToken(userID, token)
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "page":
					databaseID := util.ParseTextOne(message.Text)
					res := ctrl.svc.UpdateDatabase(userID, databaseID)
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "help":
				default:
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(util.Help)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
