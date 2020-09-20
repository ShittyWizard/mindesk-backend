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
	router.HandleFunc("/cards/{deskId}", getAllCardsByDeskId).Methods(http.MethodGet)
	router.HandleFunc("/cards/{id}", getCardById).Methods(http.MethodGet)
	router.HandleFunc("/cards", addNewCard).Methods(http.MethodPost)
	router.HandleFunc("/cards/{id}", deleteCard).Methods(http.MethodDelete)
	router.HandleFunc("/cards/{id}", editCard).Methods(http.MethodPost)
}

func getAllCards(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request getAllCards")
	err := json.NewEncoder(w).Encode(models.GetAllCards())
	if err != nil {
		log.Println(err)
	}
}

func getAllCardsByDeskId(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request getAllCards")
	deskId, _ := primitive.ObjectIDFromHex(getFieldFromRequest("deskId", r))
	err := json.NewEncoder(w).Encode(models.GetAllCardsByDeskId(deskId))
	if err != nil {
		log.Println(err)
	}
}

func getCardById(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request getCardById")
	cardId := getObjectIdFromRequest(r)
	card, err := models.GetCardById(cardId)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(&err)
	}
	err = json.NewEncoder(w).Encode(&card)
	if err != nil {
		log.Println(err)
		_ = json.NewEncoder(w).Encode(&err)
	}
}

func addNewCard(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request addNewCard")
	encoder := json.NewEncoder(w)
	var cardInitUpdate models.CardUpdate
	err := json.NewDecoder(r.Body).Decode(&cardInitUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	cardId, err := models.AddCard(cardInitUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	insertedCard, err := models.GetCardById(cardId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	err = encoder.Encode(&insertedCard)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	cardId := getObjectIdFromRequest(r)
	log.Printf("Got request deleteCard. Id: %s \n", cardId)
	err := models.DeleteCard(cardId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}

func editCard(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	cardId := getObjectIdFromRequest(r)
	log.Printf("Got request editCard. Id: %s \n", cardId)
	var cardUpdate models.CardUpdate
	err := json.NewDecoder(r.Body).Decode(&cardUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	card, err := models.EditCard(cardId, cardUpdate)
	err = json.NewEncoder(w).Encode(&card)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}

func getFieldFromRequest(field string, r *http.Request) string {
	params := mux.Vars(r)
	return params[field]
}

func getObjectIdFromRequest(r *http.Request) primitive.ObjectID {
	params := mux.Vars(r)
	cardId, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println("Can't convert 'id' to ObjectID")
		log.Println(err)
	}
	return cardId
}
