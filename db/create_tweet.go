package db

import (
	"context"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTweet(t models.CreateTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoConnection.Database("twitt")
	col := db.Collection("tweet")

	data := bson.M{
		"userId":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, data)

	if err != nil {
		return "", false, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, true, nil
}
