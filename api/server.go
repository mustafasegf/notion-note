package api

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Line *linebot.Client
	Db   *mongo.Client
}

func MakeServer(line *linebot.Client, db *mongo.Client) Server {
	server := Server{
		Line: line,
		Db:   db,
	}
	return server
}

func (s *Server) RunServer() {
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
