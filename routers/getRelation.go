package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func GetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must provide ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserId = UserID
	t.UserRelationId = ID

	var response models.GetRelation

	status, err := db.GetRelation(t)

	if err != nil || !status {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
