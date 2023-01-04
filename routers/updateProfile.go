package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Wrong data: "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := db.UpdateProfile(u, UserID)

	if err != nil {
		http.Error(w, "Was an error while updating user, try again"+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "An error ocurred while updating user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
