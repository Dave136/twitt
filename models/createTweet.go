package models

import "time"

type CreateTweet struct {
	UserId  string    `bson:"userId" json:"userId,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
