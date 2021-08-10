package service

import (
	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
)

type Line struct {
	bot    *linebot.Client
	notion core.Notion
}

func NewLinkService(bot *linebot.Client, notion core.Notion) *Line {
	return &Line{
		bot:    bot,
		notion: notion,
	}
}

func (svc *Line) CreateNote(title, body string) (status string, err error) {
	_, err = svc.notion.CreateNote(title, body)
	status = "successs"
	if err != nil {
		status = "failed"
	}
	return
}

func (svc *Line) GetLatestNote() (page *notionapi.DatabaseQueryResponse, err error) {
	page, err = svc.notion.GetLatestNote()
	return
}

func (svc *Line) AppendNote(pageID, body string) (status string, err error) {
	_, err = svc.notion.AppendNote(pageID, body)
	status = "successs"
	if err != nil {
		status = "failed"
	}
	return
}
