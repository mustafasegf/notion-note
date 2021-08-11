package service

import (
	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
	"github.com/mustafasegf/notion-note/repo"
)

type Line struct {
	bot    *linebot.Client
	notion core.Notion
	repo   *repo.Line
}

func NewLineService(bot *linebot.Client, notion core.Notion, repo *repo.Line) *Line {
	return &Line{
		bot:    bot,
		notion: notion,
		repo:   repo,
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

func (svc *Line) UpdateToken(id, token string) (status string) {
	err := svc.repo.UpdateToken(id, token)
	status = "successs"
	if err != nil {
		status = "failed"
	}
	return
}

func (svc *Line) UpdateDatabase(id, databaseID string) (status string) {
	err := svc.repo.UpdateDatabase(id, databaseID)
	status = "successs"
	if err != nil {
		status = "failed"
	}
	return
}
