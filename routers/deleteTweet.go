package routers

import (
	"net/http"

	"github.com/dave136/twitt/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must provide 'id' parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserID)

	if err != nil {
		http.Error(w, "An error ocurred while deleting tweet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
