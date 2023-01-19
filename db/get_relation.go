package db

import (
	"context"
	"fmt"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConnection.Database("twitt")
	col := database.Collection("relation")

	condition := bson.M{
		"userId":         t.UserId,
		"userRelationId": t.UserRelationId,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
