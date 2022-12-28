package jwt

import (
	"time"

	"github.com/dave136/twitt/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(u models.User) (string, error) {
	secretKey := []byte("S3cr3t-K3y")
	payload := jwt.MapClaims{
		"_id":       u.ID.Hex(),
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.LastName,
		"birthday":  u.Birthday,
		"biography": u.Biography,
		"location":  u.Location,
		"website":   u.Website,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
