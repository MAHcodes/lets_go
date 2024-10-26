package main

import (
	"log"

	"github.com/MAHcodes/lets_go/teendo/database"
	"github.com/MAHcodes/lets_go/teendo/models"
	"github.com/MAHcodes/lets_go/teendo/server"
)

func init() {
	database.Connect()
	database.Migrate(&models.Item{})
}

func main() {
	s := server.NewServer()
	log.Fatal(s.ListenAndServe())
}
