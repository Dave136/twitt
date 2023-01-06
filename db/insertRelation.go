package db

import (
	"context"
	"time"

	"github.com/dave136/twitt/models"
)

func InsertRelation(r models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConnection.Database("twitt")
	col := database.Collection("relation")

	_, err := col.InsertOne(ctx, r)

	if err != nil {
		return false, err
	}

	return true, nil
}
