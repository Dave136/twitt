package routers

import (
	"errors"
	"strings"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
	"github.com/golang-jwt/jwt/v4"
)

var Email string
var UserID string

func TokenProcess(token string) (*models.Claim, bool, string, error) {
	secretKey := []byte("S3cr3t-K3y")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	newToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err == nil {
		_, found, ID := db.CheckUserExist(claims.Email)

		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}

		return claims, found, ID, nil
	}

	if !newToken.Valid {
		return claims, false, "", errors.New("invalid token")
	}

	return claims, false, "", err
}
