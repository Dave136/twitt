package db

import (
	"context"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTweetsFollowers(ID string, page int) ([]models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConnection.Database("twitt")
	col := database.Collection("relation")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{
		"userId": ID,
	}})

	// Join 2 documents
	// https://www.mongodb.com/docs/v6.0/reference/operator/aggregation/lookup/#examples
	conditions = append(conditions, bson.M{"$lookup": bson.M{
		"from":         "tweet",
		"localField":   "userRelationId",
		"foreignField": "userId",
		"as":           "tweet",
	}})

	// https://www.mongodb.com/docs/manual/reference/operator/aggregation/unwind/#examples
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.TweetsFollowers

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}
