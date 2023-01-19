package middlewares

import (
	"net/http"

	"github.com/dave136/twitt/db"
)

func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost connection with the database", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
