package api

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
)

type Server struct {
	Line   *linebot.Client
	Notion core.Notion
}

func MakeServer(line *linebot.Client, notion core.Notion) Server {
	server := Server{
		Line:   line,
		Notion: notion,
	}
	return server
}

func (s *Server) RunServer() {
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
