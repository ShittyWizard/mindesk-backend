package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"min-desk-backend/internal/models"
	"net/http"
)

func initCardsHandlers() {
	router.HandleFunc("/cards", getAllCards).Methods(http.MethodGet)
	router.HandleFunc("/cards", addNewCard).Methods(http.MethodPost)
	router.HandleFunc("/cards/{id}", deleteCard).Methods(http.MethodDelete)
	router.HandleFunc("/cards/{id}", editCard).Methods(http.MethodPost)
}

func getAllCards(w http.ResponseWriter, r *http.Request) {
	log.Println("Got response getAllCards")
	err := json.NewEncoder(w).Encode(models.GetAllCards())
	if err != nil {
		log.Fatal(err)
	}
}

func addNewCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Got response addNewCard")
	var card models.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		log.Fatal(err)
	}
	models.AddCard(card)
	err = json.NewEncoder(w).Encode(&card)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	cardId := getCardObjectId(r)
	log.Printf("Got response deleteCard. Id: %s \n", cardId)
	err := models.DeleteCard(cardId)
	if err != nil {
		log.Fatal(err)
	}
}

func editCard(w http.ResponseWriter, r *http.Request) {
	cardId := getCardObjectId(r)
	log.Printf("Got response editCard. Id: %s \n", cardId)
	var cardUpdate models.CardUpdate
	err := json.NewDecoder(r.Body).Decode(&cardUpdate)
	if err != nil {
		log.Fatal(err)
	}
	card := models.EditCard(cardId, cardUpdate)
	err = json.NewEncoder(w).Encode(&card)
	if err != nil {
		log.Fatal(err)
	}
}

func getCardObjectId(r *http.Request) primitive.ObjectID {
	params := mux.Vars(r)
	cardId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println("Can't convert 'id' to ObjectID")
		log.Fatal(err)
	}
	return cardId
}
