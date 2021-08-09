package main

import (
	"github.com/jomei/notionapi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/api"
	"github.com/mustafasegf/notion-note/core"
	"github.com/mustafasegf/notion-note/util"
)

func main() {
	config, _ := util.LoadConfig()

	client := notionapi.NewClient(notionapi.Token(config.NotionToken))
	bot, _ := linebot.New(config.LineSecret, config.LineToken)

	notion := core.Notion{
		Config: config,
		Client: client,
	}

	server := api.MakeServer(bot, notion)
	server.SetupRouter()
	server.RunServer()
}
