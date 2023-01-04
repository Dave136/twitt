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

	if PORT == "" {
		PORT = "1173"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
