package main

import (
	"log"
	"min-desk-backend/internal/handlers"
	"min-desk-backend/internal/models"
	"net/http"
)

// Mongo address in docker-compose
var mongoAddress = "mongodb://mongodb:27017"

func main() {
	log.Println("Server started...")
	models.InitMongoDB(mongoAddress)
	models.InitCardsCollection()

	// adding test data
	models.AddTestCards()
	log.Println("Test data was added")

	r := handlers.InitRouter()
	log.Println("Router was initialised")
	log.Fatal(http.ListenAndServe(":8080", r))
}
