package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/dave136/twitt/db"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must provide ID", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(w, "Image not found"+err.Error(), http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error copying image"+err.Error(), http.StatusNotFound)
	}
}
