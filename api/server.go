package api

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Line   *linebot.Client
	Notion core.Notion
	Db     *mongo.Client
}

func MakeServer(line *linebot.Client, notion core.Notion, db *mongo.Client) Server {
	server := Server{
		Line:   line,
		Notion: notion,
		Db:     db,
	}
	return server
}

func (s *Server) RunServer() {
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
