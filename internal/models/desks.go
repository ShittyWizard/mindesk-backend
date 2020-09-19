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
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty"`
}

type DeskUpdate struct {
	Name string `json:"name,omitempty"`
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

	_, err := desksCollection.InsertOne(context.Background(), desk)
	if err != nil {
		log.Println(err)
	}

	return deskId, err
}

func GetTestDesk() (Desk, error) {
	var desk Desk
	err := desksCollection.FindOne(context.Background(), bson.D{}).Decode(&desk)
	if err != nil {
		log.Println(err)
	}
	return desk, err
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

func DeleteDesk(deskId primitive.ObjectID) error {
	_, err := desksCollection.DeleteOne(context.Background(), bson.D{{"_id", deskId}})
	if err != nil {
		log.Println(err)
	}
	return err
}
