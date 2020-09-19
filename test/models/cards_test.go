package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"min-desk-backend/internal/models"
	"testing"
)

var insertedCardId primitive.ObjectID

func TestSetupDevEnv(t *testing.T) {
	models.InitMongoDB(models.TestMongoAddress)
	models.InitCardsCollection()

	// Add 4 test cards
	models.AddTestCards()
}

func TestGetAllCards(t *testing.T) {
	expectedSize := 4
	realSize := len(models.GetAllCards())
	if expectedSize != realSize {
		t.Errorf("Get all cards works incorrectly. Expected size: %d, real size: %d", expectedSize, realSize)
	}
}

func TestAddCard(t *testing.T) {
	initSize := len(models.GetAllCards())
	insertedCardId = models.AddCard(models.CardUpdate{
		Name:        "test1",
		Description: "test2 description",
		AssignedTo:  "John Johnson",
		DueDate:     "31-12-2020",
	})
	newSize := len(models.GetAllCards())
	if newSize <= initSize {
		t.Errorf("Adding card works incorrectly. Init size: %d, new size: %d", initSize, newSize)
	}
}
func TestEditCard(t *testing.T) {
	cardUpdate := models.CardUpdate{
		Name:        "Updated name",
		Description: "Updated description",
		AssignedTo:  "Peter Peterson",
		DueDate:     "25-09-2020",
	}
	editCard := models.EditCard(insertedCardId, cardUpdate)
	if len(cardUpdate.Name) != 0 && editCard.Name != cardUpdate.Name {
		t.Errorf("Editing of card works incorrectly. Card's name: %s, CardUpdate's name: %s", editCard.Name, cardUpdate.Name)
	}
	if len(cardUpdate.Description) != 0 && editCard.Description != cardUpdate.Description {
		t.Errorf("Editing of card works incorrectly. Card's Description: %s, CardUpdate's Description: %s", editCard.Description, cardUpdate.Description)
	}
	if len(cardUpdate.AssignedTo) != 0 && editCard.AssignedTo != cardUpdate.AssignedTo {
		t.Errorf("Editing of card works incorrectly. Card's AssignedTo: %s, CardUpdate's AssignedTo: %s", editCard.AssignedTo, cardUpdate.AssignedTo)
	}
	if len(cardUpdate.DueDate) != 0 && editCard.DueDate.Time().Format("02-01-2006") != cardUpdate.DueDate {
		t.Errorf("Editing of card works incorrectly. Card's DueDate: %T, CardUpdate's DueDate: %T", &editCard.DueDate, &cardUpdate.DueDate)
	}
}

func TestDeleteCard(t *testing.T) {
	sizeBeforeDeleting := len(models.GetAllCards())
	err := models.DeleteCard(insertedCardId)
	if err != nil {
		log.Fatal(err)
	}
	sizeAfterDeleting := len(models.GetAllCards())
	if sizeAfterDeleting >= sizeBeforeDeleting {
		t.Errorf("Deleting card works incorrectly, Size before deleting: %d, size after deleting: %d", sizeBeforeDeleting, sizeAfterDeleting)
	}
}
