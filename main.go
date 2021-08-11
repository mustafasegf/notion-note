package main

import (
	"context"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/notion-note/api"
	"github.com/mustafasegf/notion-note/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	bot, err := linebot.New(config.LineSecret, config.LineToken)
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}

	util.SetLogger()
	server := api.MakeServer(bot, db)
	server.SetupRouter()
	server.RunServer(config.ServerPort)

	defer func() {
		if err = db.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
