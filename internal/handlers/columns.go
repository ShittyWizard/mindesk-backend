package handlers

import (
	"encoding/json"
	"log"
	"min-desk-backend/internal/models"
	"net/http"
)

func initColumnsHandlers() {
	router.HandleFunc("/columns", addColumn).Methods(http.MethodPost)
	router.HandleFunc("/columns/{id}", deleteColumn).Methods(http.MethodDelete)
}

func addColumn(w http.ResponseWriter, r *http.Request) {
	log.Println("Got request addColumn")
	encoder := json.NewEncoder(w)
	var columnInitUpdate models.ColumnUpdate
	err := json.NewDecoder(r.Body).Decode(&columnInitUpdate)

	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	columnId, err := models.AddColumnToDesk(columnInitUpdate)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
	err = encoder.Encode(&columnId)
}

func deleteColumn(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	columnId := getObjectIdFromRequest(r)
	log.Printf("Got request deleteColumn. Id: %s \n", columnId)
	err := models.DeleteColumnFromDesk(columnId)
	if err != nil {
		log.Println(err)
		_ = encoder.Encode(&err)
	}
}
