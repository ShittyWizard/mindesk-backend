package handlers

import (
	"encoding/json"
	"log"
	"min-desk-backend/internal/models"
	"net/http"
)

func initDesksHandlers() {
	router.HandleFunc("/desks", getTestDesk).Methods(http.MethodGet)
	router.HandleFunc("/desks/{id}", getDeskById).Methods(http.MethodGet)
	router.HandleFunc("/desks", addDesk).Methods(http.MethodPost)
	router.HandleFunc("/desks/{id}", deleteDesk).Methods(http.MethodDelete)
	router.HandleFunc("/desks/{id}", editDesk).Methods(http.MethodPost)
}

func getTestDesk(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request getTestDesk")
	encoder := json.NewEncoder(w)
	desk, err := models.GetTestDesk()
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(err)
	}
	err = encoder.Encode(desk)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(err)
	}
}

func getDeskById(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request getDeskById")
	encoder := json.NewEncoder(w)
	deskId := getObjectIdFromRequest(r)
	desk, err := models.GetDeskById(deskId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(err)
	}
	err = encoder.Encode(desk)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(err)
	}
}

func addDesk(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request addDesk")
	encoder := json.NewEncoder(w)
	var deskInitUpdate models.DeskUpdate
	err := json.NewDecoder(r.Body).Decode(&deskInitUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	deskId, err := models.AddDesk(deskInitUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	insertedDesk, err := models.GetDeskById(deskId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	err = encoder.Encode(&insertedDesk)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}

func deleteDesk(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	deskId := getObjectIdFromRequest(r)
	log.Printf("Got request deleteDesk. Id: %s \n", deskId)
	err := models.DeleteDesk(deskId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}

func editDesk(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	deskId := getObjectIdFromRequest(r)
	log.Printf("Got request editDesk. Id: %s \n", deskId)
	var deskUpdate models.DeskUpdate
	err := json.NewDecoder(r.Body).Decode(&deskUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	desk, err := models.EditDesk(deskId, deskUpdate)
	err = json.NewEncoder(w).Encode(&desk)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}
