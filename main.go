package main

import (
	"log"

	"github.com/dave136/twitt/db"
	"github.com/dave136/twitt/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Cannot connect to database")
		return
	}

	handlers.Handlers()
}
