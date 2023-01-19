package db

import (
	"context"
	"log"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTwets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoConnection.Database("twitt")
	col := db.Collection("tweet")

	var results []*models.GetTwets
	condition := bson.M{
		"userId": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.GetTwets
		err = cursor.Decode(&tweet)
		if err != nil {
			return results, false
		}
		results = append(results, &tweet)
	}

	return results, true
}
