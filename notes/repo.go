package notes

import (
	"context"
	"time"

	"github.com/mustafasegf/notion-note/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	DB *mongo.Client
}

func NewRepo(db *mongo.Client) *Repo {
	return &Repo{
		DB: db,
	}
}

func (repo *Repo) UpdateToken(id, token string) (err error) {
	coll := repo.DB.Database("users").Collection("line")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: token}}}}

	_, err = coll.UpdateOne(ctx, filter, update, opts)
	return
}

func (repo *Repo) UpdateDatabase(id, databaseID string) (err error) {
	coll := repo.DB.Database("users").Collection("line")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "databaseID", Value: databaseID}}}}

	_, err = coll.UpdateOne(ctx, filter, update, opts)
	return
}

func (repo *Repo) GetNotionCreds(id string) (res entity.NotionCreds, err error) {
	coll := repo.DB.Database("users").Collection("line")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: id}}
	err = coll.FindOne(ctx, filter).Decode(&res)
	return
}
