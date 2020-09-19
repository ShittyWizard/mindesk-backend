package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"min-desk-backend/internal/models"
	"testing"
)

var insertedDeskId primitive.ObjectID
var testDeskId primitive.ObjectID

func TestDesksSetupDevEnv(t *testing.T) {
	models.InitMongoDB(models.TestMongoAddress)

	testDeskId = models.AddTestDesks()[0]
}

func TestGetDeskById(t *testing.T) {
	desk, err := models.GetDeskById(testDeskId)
	if err != nil || len(desk.Name) == 0 {
		t.Errorf("Getting desk by id works incorrectly...")
	}
}

func TestAddDesk(t *testing.T) {
	var err error
	insertedDeskId, err = models.AddDesk(models.DeskUpdate{Name: "Test desk 1"})
	if err != nil {
		t.Errorf("Adding desk works incorrectly...")
	}
	insertedDesk, err := models.GetDeskById(insertedDeskId)
	if err != nil || len(insertedDesk.Name) == 0 {
		t.Errorf("Adding desk works incorrectly...")
	}
}

func TestEditDesk(t *testing.T) {
	deskUpdate := models.DeskUpdate{
		Name: "Name was updated",
	}
	editDesk, err := models.EditDesk(insertedDeskId, deskUpdate)
	if err != nil || (len(deskUpdate.Name) != 0 && editDesk.Name != deskUpdate.Name) {
		t.Errorf("Editing of desk works incorrectly. desk's name: %s, deskUpdate's name: %s", editDesk.Name, deskUpdate.Name)
	}
}

func TestDeleteDesk(t *testing.T) {
	err := models.DeleteDesk(insertedDeskId)
	if err != nil {
		t.Errorf("Deleting desk works incorrectly...")
	}
	_, err = models.GetDeskById(insertedDeskId)
	if err == nil {
		t.Errorf("Deleting desk works incorrectly...")
	}
}
