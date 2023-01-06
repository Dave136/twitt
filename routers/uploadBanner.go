package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var ext = strings.Split(handler.Filename, ".")[1]
	var imgName = UserID + "." + ext
	var filepath string = "uploads/banners/" + imgName

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "An error ocurred while uploading banner: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "An error copying the image: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = imgName
	status, err = db.UpdateProfile(user, UserID)

	if err != nil || !status {
		http.Error(w, "An error ocurred while uploading banner: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
