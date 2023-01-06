package routers

import (
	"net/http"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must provide ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserId = UserID
	t.UserRelationId = ID

	status, err := db.DeleteRelation(t)

	if err != nil || !status {
		http.Error(w, "An error ocurred while deleting relation"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
