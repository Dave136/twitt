package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "You need to provide a message", http.StatusBadRequest)
		return
	}

	data := models.CreateTweet{
		UserId:  UserID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.CreateTweet(data)

	if err != nil {
		http.Error(w, "An error ocurred while trying creating a new tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "An error ocurred while trying creating a tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
