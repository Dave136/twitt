package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dave136/twitt/middlewares"
	"github.com/dave136/twitt/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")

	router.HandleFunc("/register", middlewares.CheckDatabase(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDatabase(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDatabase(middlewares.JwtValidate(routers.Profile))).Methods("GET")
	router.HandleFunc("/profile", middlewares.CheckDatabase(middlewares.JwtValidate(routers.UpdateProfile))).Methods("PUT")

	// Tweets
	router.HandleFunc("/tweet", middlewares.CheckDatabase(middlewares.JwtValidate(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlewares.CheckDatabase(middlewares.JwtValidate(routers.CreateTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlewares.CheckDatabase(middlewares.JwtValidate(routers.DeleteTweet))).Methods("DELETE")

	// Images
	router.HandleFunc("/avatar", middlewares.CheckDatabase(middlewares.JwtValidate(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlewares.CheckDatabase(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/banner", middlewares.CheckDatabase(middlewares.JwtValidate(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/banner", middlewares.CheckDatabase(routers.GetBanner)).Methods("GET")

	// Relation
	router.HandleFunc("/relation", middlewares.CheckDatabase(middlewares.JwtValidate(routers.GetRelation))).Methods("GET")
	router.HandleFunc("/relation", middlewares.CheckDatabase(middlewares.JwtValidate(routers.Relation))).Methods("POST")
	router.HandleFunc("/relation", middlewares.CheckDatabase(middlewares.JwtValidate(routers.DeleteRelation))).Methods("DELETE")

	router.HandleFunc("/following", middlewares.CheckDatabase(middlewares.JwtValidate(routers.GetUsers))).Methods("GET")

	if PORT == "" {
		PORT = "1173"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
