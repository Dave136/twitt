package db

import (
	"context"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("twitt")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}

	register["birthday"] = u.Birthday

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.Website) > 0 {
		register["website"] = u.Website
	}

	updateString := bson.M{
		"$set": register,
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
