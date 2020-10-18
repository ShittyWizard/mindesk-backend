package main

import (
	"github.com/gorilla/handlers"
	"log"
	"min-desk-backend/internal/api_handlers"
	"min-desk-backend/internal/models"
	"net/http"
)

func main() {
	log.Println("Server started...")
	models.InitMongoDB(models.MongoAddress)

	log.Println("Test data was added")
	r := api_handlers.InitRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Println("Router was initialised")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
