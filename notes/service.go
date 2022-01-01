package notes

import (
	"fmt"
	"log"
	"strings"

	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	bot  *linebot.Client
	repo *Repo
}

func NewService(bot *linebot.Client, repo *Repo) *Service {
	return &Service{
		bot:  bot,
		repo: repo,
	}
}

func (svc *Service) CreateNote(userID, title, body, tags string) (status string, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		status = err.Error()
		return
	}
	notion := HttpRequest{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	if tags == "" {
		_, err = notion.CreateNote(title, body, creds.DatabaseID)
	} else {
		tagsArr := strings.Split(tags, " ")
		_, err = notion.CreateNoteWithTags(title, body, tagsArr, creds.DatabaseID)
	}

	status = "successs"
	if err != nil {
		status = err.Error()
	}
	return
}

func (svc *Service) GetLatestNote(userID string) (page *notionapi.DatabaseQueryResponse, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		return
	}
	notion := HttpRequest{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	page, err = notion.GetLatestNote(creds.DatabaseID)
	return
}

func (svc *Service) FindNote(userID, query string) (page *notionapi.DatabaseQueryResponse, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		log.Println(err)
		return
	}
	notion := HttpRequest{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	page, err = notion.FindNote(creds.DatabaseID, query)
	return
}

func (svc *Service) AppendNote(userID, pageID, body string) (status string, err error) {
	creds, err := svc.GetNotionCreds(userID)
	if err != nil {
		status = err.Error()
		return
	}
	notion := HttpRequest{Client: notionapi.NewClient(notionapi.Token(creds.Token))}
	_, err = notion.AppendNote(pageID, body)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Service) UpdateToken(userID, token string) (status string) {
	err := svc.repo.UpdateToken(userID, token)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Service) UpdateDatabase(userID, databaseID string) (status string) {
	err := svc.repo.UpdateDatabase(userID, databaseID)
	status = "successs"
	if err != nil {
		log.Println(err)
		status = "failed"
	}
	return
}

func (svc *Service) GetNotionCreds(userID string) (res entity.NotionCreds, err error) {
	res, err = svc.repo.GetNotionCreds(userID)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Println(err)
		return
	}

	if res.DatabaseID == "" {
		err = fmt.Errorf("no database id. Please use /page [databaseID] to set database\n")
	}
	if res.Token == "" {
		err = fmt.Errorf("no notion secret token. Go to https://www.notion.so/my-integrations to get your token and use /token [token] to set token\n\n%v", err)
	}
	return
}
