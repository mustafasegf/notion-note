package main

import (
	"github.com/jomei/notionapi"
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

	notion.CreateNote("ini title", "ini content \nyang keren")

}
