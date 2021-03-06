package main

import (
	"log"
	"min-desk-backend/internal/handlers"
	"min-desk-backend/internal/models"
	"net/http"
)

func main() {
	log.Println("Server started...")
	models.InitMongoDB(models.MongoAddress)

	log.Println("Test data was added")

	r := handlers.InitRouter()
	log.Println("Router was initialised")
	log.Fatal(http.ListenAndServe(":8080", r))
}
