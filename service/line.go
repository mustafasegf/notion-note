package service

import (
	"fmt"
	"log"

	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
	"github.com/mustafasegf/notion-note/entity"
	"github.com/mustafasegf/notion-note/repo"
)

type Line struct {
	bot  *linebot.Client
	repo *repo.Line
}

func NewLineService(bot *linebot.Client, repo *repo.Line) *Line {
	return &Line{
		bot:  bot,
		repo: repo,
	}
}

func (svc *Line) CreateNote(userID, title, body string) (status string, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		status = "failed"
		return
	}
	notion := core.Notion{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	_, err = notion.CreateNote(title, body, creds.DatabaseID)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Line) GetLatestNote(userID string) (page *notionapi.DatabaseQueryResponse, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		return
	}
	notion := core.Notion{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	page, err = notion.GetLatestNote(creds.DatabaseID)
	return
}

func (svc *Line) AppendNote(userID, pageID, body string) (status string, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		status = "failed"
		return
	}
	notion := core.Notion{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	_, err = notion.AppendNote(pageID, body)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Line) UpdateToken(userID, token string) (status string) {
	err := svc.repo.UpdateToken(userID, token)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Line) UpdateDatabase(userID, databaseID string) (status string) {
	err := svc.repo.UpdateDatabase(userID, databaseID)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Line) GetNotionCreds(userID string) (res entity.NotionCreds, err error) {
	res, err = svc.repo.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		return
	}
	if res.DatabaseID == "" {
		err = fmt.Errorf("no database id")
	}
	if res.Token == "" {
		err = fmt.Errorf("%v, no token", err)
	}
	return
}
