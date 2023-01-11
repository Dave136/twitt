package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"userId" json:"userId,omitempty"`
	UserRelationId string             `bson:"userRelationId" json:"userRelationId,omitempty"`
	Tweet          struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
	}
}
