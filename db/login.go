package db

import (
	"github.com/dave136/twitt/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExist(email)

	if !found {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
