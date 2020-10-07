package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"strconv"
)

type Desk struct {
	Id        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name,omitempty"`
	ColumnIds []primitive.ObjectID `json:"columnIds,omitempty" bson:"columnIds"`
}

type DeskUpdate struct {
	Name     string `json:"name,omitempty"`
	ColumnId string `json:"columndId,omitempty"`
}

var desksCollectionName = "desks"
var desksCollection *mongo.Collection

func InitDesksCollection() {
	log.Println("Initialising DESKS collection...")
	desksCollection = mongodb.Collection(desksCollectionName)
	log.Println("DESKS collection initialised...")
}

func AddTestDesks() []primitive.ObjectID {
	desksCollection.DeleteMany(context.Background(), bson.D{})
	testDesk := DeskUpdate{Name: "Gachi Dungeon Gym #" + strconv.Itoa(rand.Int())}
	deskId, _ := AddDesk(testDesk)
	return []primitive.ObjectID{deskId}
}

func AddDesk(deskUpdate DeskUpdate) (primitive.ObjectID, error) {
	deskId := primitive.NewObjectID()
	var desk Desk
	desk.Id = deskId
	desk.Name = deskUpdate.Name
	desk.ColumnIds = []primitive.ObjectID{primitive.NewObjectID()}
	desk.ColumnIds = append(desk.ColumnIds[:0], desk.ColumnIds[1:]...)

	_, err := desksCollection.InsertOne(context.Background(), desk)
	if err != nil {
		log.Println(err)
	}

	return deskId, err
}

func GetAllDesks() ([]*Desk, error) {
	var desks []*Desk
	cursor, err := desksCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		desk := Desk{}
		err := cursor.Decode(&desk)
		if err != nil {
			log.Println(err)
		}
		desks = append(desks, &desk)
	}

	return desks, err
}

func GetDeskById(deskId primitive.ObjectID) (Desk, error) {
	var desk Desk
	err := desksCollection.FindOne(context.Background(), bson.D{{"_id", deskId}}).Decode(&desk)
	if err != nil {
		log.Println(err)
	}
	return desk, err
}

func EditDesk(deskId primitive.ObjectID, deskUpdate DeskUpdate) (Desk, error) {
	updatedFields := bson.D{}
	if len(deskUpdate.Name) != 0 {
		updatedFields = append(updatedFields, bson.E{Key: "name", Value: deskUpdate.Name})
	}
	if len(updatedFields) == 0 {
		return GetDeskById(deskId)
	}
	_, err := desksCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": deskId},
		bson.D{
			{"$set", updatedFields},
		})
	if err != nil {
		log.Println(err)
		return Desk{}, err
	}
	desk, err := GetDeskById(deskId)
	if err != nil {
		log.Println(err)
		return Desk{}, err
	} else {
		return desk, err
	}
}

func UpdateDeskColumns(deskId primitive.ObjectID, columnId primitive.ObjectID, action string) error {
	var err error
	if action == ColumnAddAction {
		err = addColumnIdToDesk(columnId, deskId)
	} else if action == ColumnRemoveAction {
		err = removeColumnIdFromDesk(columnId, deskId)
	}
	return err
}

func DeleteDesk(deskId primitive.ObjectID) error {
	_, err := desksCollection.DeleteOne(context.Background(), bson.D{{"_id", deskId}})
	if err != nil {
		log.Println(err)
	}
	return err
}

func addColumnIdToDesk(columnId primitive.ObjectID, deskId primitive.ObjectID) error {
	log.Println("Add column " + columnId.Hex() + " to desk " + deskId.Hex())
	_, err := desksCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": deskId},
		bson.D{
			{"$push", bson.M{"columnIds": columnId}},
		})
	return err
}

func removeColumnIdFromDesk(columnId primitive.ObjectID, deskId primitive.ObjectID) error {
	log.Println("Add column " + columnId.Hex() + " to desk " + deskId.Hex())
	_, err := desksCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": deskId},
		bson.D{
			{"$pull", bson.M{"columnIds": columnId}},
		})
	return err
}
