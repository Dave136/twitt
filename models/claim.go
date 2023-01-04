package models

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
