package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dave136/twitt/db"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	tempPage, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "You must provide a valid page number, it need to be greather than 0", http.StatusBadRequest)
		return
	}

	pag := int64(tempPage)

	result, status := db.GetAllUsers(UserID, pag, search, typeUser)

	if !status {
		http.Error(w, "An error ocurred while getting users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
