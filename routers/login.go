package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/jwt"
	"github.com/dave136/twitt/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "User or password incorrect: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	document, exists := db.Login(u.Email, u.Password)

	if !exists {
		http.Error(w, "User or password incorrect", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Was an error while login"+err.Error(), http.StatusInternalServerError)
		return
	}

	result := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

	expTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expTime,
		// Path:     "/",
		// HttpOnly: true,
	})
}
