package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"min-desk-backend/internal/models"
	"net/http"
)

func initCardsHandlers() {
	router.HandleFunc("/cards", getAllCards).Methods(http.MethodGet)
	router.HandleFunc("/cards", addNewCard).Methods(http.MethodPost)
	router.HandleFunc("/cards/{id}", deleteCard).Methods(http.MethodDelete)
}

func getAllCards(w http.ResponseWriter, r *http.Request) {
	log.Println("Got response GetAllCards")
	err := json.NewEncoder(w).Encode(models.AllCards())
	if err != nil {
		log.Fatal(err)
	}
}

func addNewCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	_ = json.NewDecoder(r.Body).Decode(&card)
	models.AddCard(card)
	err := json.NewEncoder(w).Encode(&card)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := models.DeleteCard(params["id"])
	if err != nil {
		log.Fatal(err)
	}
}
