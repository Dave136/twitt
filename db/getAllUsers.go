package db

import (
	"context"
	"fmt"
	"time"

	"github.com/dave136/twitt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConnection.Database("twitt")
	col := database.Collection("users")

	var results []*models.User

	options := options.Find()
	options.SetSkip((page - 1) * 20)
	options.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, options)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserId = ID
		r.UserRelationId = s.ID.Hex()

		include = false

		found, err = GetRelation(r)

		if userType == "new" && !found {
			include = true
		}

		if userType == "follow" && found {
			include = true
		}

		if r.UserRelationId == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.Banner = ""
			s.Email = ""
			s.Location = ""
			s.Website = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)

	return results, true
}
