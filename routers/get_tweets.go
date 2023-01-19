package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dave136/twitt/db"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must provide the 'id' parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must provide the 'page' parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "You must provide the 'page' parameter greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	result, ok := db.GetTweets(ID, pag)

	if !ok {
		http.Error(w, "An error ocurred while getting tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
