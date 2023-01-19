package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dave136/twitt/db"
)

func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must provide a page parameter", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "You must provide a page parameter, must be greater than 0", http.StatusBadRequest)
		return
	}

	result, right := db.GetTweetsFollowers(UserID, pag)

	if !right {
		http.Error(w, "An error ocurred while getting tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
