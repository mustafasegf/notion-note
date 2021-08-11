package repo

import (
	"context"
	"time"

	"github.com/mustafasegf/notion-note/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Line struct {
	Db *mongo.Client
}

func NewLineRepo(db *mongo.Client) *Line {
	return &Line{
		Db: db,
	}
}

func (repo *Line) UpdateToken(id, token string) (err error) {
	coll := repo.Db.Database("users").Collection("line")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"token", token}}}}

	_, err = coll.UpdateOne(ctx, filter, update, opts)
	return
}

func (repo *Line) UpdateDatabase(id, databaseID string) (err error) {
	coll := repo.Db.Database("users").Collection("line")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"databaseID", databaseID}}}}

	_, err = coll.UpdateOne(ctx, filter, update, opts)
	return
}

func (repo *Line) GetNotionCreds(id string) (res entity.NotionCreds, err error) {
	coll := repo.Db.Database("users").Collection("line")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	filter := bson.D{{"_id", id}}
	err = coll.FindOne(ctx, filter).Decode(&res)
	return
}
