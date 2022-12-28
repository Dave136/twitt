package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "An error has occurred: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "Password must be at less 6 characters", http.StatusBadRequest)
		return
	}

	_, found, _ := db.CheckUserExist(u.Email)

	if found {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertIntoDatabase(u)

	if err != nil {
		http.Error(w, "An error ocurred while creating user: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Cannot create new user: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
