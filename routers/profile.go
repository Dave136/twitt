package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dave136/twitt/db"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Must send the 'id' parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "An error ocurred while getting profile", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
