package db

import (
	"context"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	db := MongoConnection.Database("twitt")
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": objID}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		return profile, err
	}

	return profile, nil
}
