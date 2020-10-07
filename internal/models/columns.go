package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Column struct {
	Id      primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string               `json:"name,omitempty"`
	CardIds []primitive.ObjectID `json:"cardIds,omitempty" bson:"cardIds"`
	DeskId  primitive.ObjectID   `json:"deskId,omitempty" bson:"deskId,omitempty"`
}

type ColumnUpdate struct {
	Name   string `json:"name,omitempty"`
	DeskId string `json:"deskId,omitempty"`
}

var columnsCollectionName = "columns"
var columnsCollection *mongo.Collection

var ColumnAddAction = "COLUMN_ADD"
var ColumnRemoveAction = "COLUMN_REMOVE"

func InitColumnsCollection() {
	log.Println("Initialising COLUMNS collection...")
	columnsCollection = mongodb.Collection(columnsCollectionName)
	log.Println("COLUMNS collection initialised...")
}

func AddTestColumns(deskId primitive.ObjectID) primitive.ObjectID {
	columnsCollection.DeleteMany(context.Background(), bson.D{})
	id, _ := AddColumnToDesk(ColumnUpdate{
		Name:   "To do",
		DeskId: deskId.Hex(),
	})
	return id
}

func getColumnById(id primitive.ObjectID) (Column, error) {
	var column Column
	err := columnsCollection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&column)
	if err != nil {
		log.Println("Error with getting column by id: " + id.String())
	}
	return column, err
}

func GetColumnsByDeskId(deskId primitive.ObjectID) []Column {
	desk, _ := GetDeskById(deskId)
	var columns []Column
	for _, columnId := range desk.ColumnIds {
		column, _ := getColumnById(columnId)
		columns = append(columns, column)
	}

	return columns
}

func AddColumnToDesk(update ColumnUpdate) (primitive.ObjectID, error) {
	deskId, _ := primitive.ObjectIDFromHex(update.DeskId)
	columnId, _ := addColumnToCollection(update.Name, deskId)
	err := UpdateDeskColumns(deskId, columnId, ColumnAddAction)
	return columnId, err
}

func addColumnToCollection(name string, deskId primitive.ObjectID) (primitive.ObjectID, error) {
	columnId := primitive.NewObjectID()
	var column Column
	column.Id = columnId
	column.Name = name
	column.DeskId = deskId
	column.CardIds = []primitive.ObjectID{primitive.NewObjectID()}
	column.CardIds = append(column.CardIds[:0], column.CardIds[1:]...)

	_, err := columnsCollection.InsertOne(context.Background(), column)
	if err != nil {
		log.Println(err)
	}

	return columnId, err
}

func DeleteColumnFromDesk(columnId primitive.ObjectID) error {
	column, _ := getColumnById(columnId)
	for _, cardId := range column.CardIds {
		_ = DeleteCard(cardId)
	}
	deskId := column.DeskId
	_ = deleteColumnFromCollection(columnId)
	err := UpdateDeskColumns(deskId, columnId, ColumnRemoveAction)
	return err
}

func deleteColumnFromCollection(columnId primitive.ObjectID) error {
	_, err := columnsCollection.DeleteOne(context.Background(), bson.D{{"_id", columnId}})
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateColumnCards(columnId primitive.ObjectID, cardId primitive.ObjectID, action string) error {
	var err error
	log.Println("Update column cards " + columnId.Hex() + " " + cardId.Hex() + " " + action)
	if action == CardAddAction {
		err = addCardIdToColumn(cardId, columnId)
	} else if action == CardRemoveAction {
		err = removeCardIdFromColumn(cardId, columnId)
	}
	return err
}

func addCardIdToColumn(cardId primitive.ObjectID, columnId primitive.ObjectID) error {
	log.Println("Add card " + cardId.Hex() + " to column " + columnId.Hex())
	_, err := columnsCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": columnId},
		bson.D{
			{"$push", bson.M{"cardIds": cardId}},
		})
	return err
}

func removeCardIdFromColumn(cardId primitive.ObjectID, columnId primitive.ObjectID) error {
	log.Println("Add card " + cardId.Hex() + " to column " + columnId.Hex())
	_, err := columnsCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": columnId},
		bson.D{
			{"$pull", bson.M{"cardIds": cardId}},
		})
	return err
}
