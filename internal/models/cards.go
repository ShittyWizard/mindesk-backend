package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Card struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	AssignedTo  string             `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
	DueDate     primitive.DateTime `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
}

type CardUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AssignedTo  string `json:"assignedTo,omitempty" bson:"assignedTo,omitempty"`
	DueDate     string `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
}

var cardsCollectionName = "cards"
var cardsCollection *mongo.Collection

func InitCardsCollection() {
	log.Println("Trying to connect to MongoDB...")
	cardsCollection = mongodb.Collection(cardsCollectionName)
	log.Println("Connected to MongoDB...")
}

// If you change size of test cards - test it too in test/models/cards_test.go:17
func AddTestCards() {
	cardsCollection.DeleteMany(context.Background(), bson.D{})
	testCard1 := CardUpdate{"test name 1", "test description 1", "me1", "25-09-2020"}
	AddCard(testCard1)

	testCard2 := CardUpdate{"test name 2", "test description 2", "me2", "31-12-2020"}
	AddCard(testCard2)

	testCard3 := CardUpdate{"test name 3", "test description 3", "me3", "28-10-2020"}
	AddCard(testCard3)

	testCard4 := CardUpdate{"test name 4", "test description 4", "me4", "26-05-2021"}
	AddCard(testCard4)
}

func GetAllCards() []*Card {
	var cards []*Card
	cursor, err := cardsCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		card := Card{}
		err := cursor.Decode(&card)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, &card)
	}

	return cards
}

func GetCardById(cardId primitive.ObjectID) Card {
	var card Card
	err := cardsCollection.FindOne(context.Background(), bson.D{{"_id", cardId}}).Decode(&card)
	if err != nil {
		log.Fatal(err)
	}
	return card
}

func AddCard(cardUpdate CardUpdate) primitive.ObjectID {
	cardId := primitive.NewObjectID()
	var card Card
	card.Id = cardId
	card.Name = cardUpdate.Name
	card.Description = cardUpdate.Description
	card.AssignedTo = cardUpdate.AssignedTo

	dueDate, err := time.Parse("02-01-2006", cardUpdate.DueDate)
	if err != nil {
		log.Fatal(err)
	}
	card.DueDate = primitive.NewDateTimeFromTime(dueDate)
	_, err = cardsCollection.InsertOne(context.Background(), card)
	if err != nil {
		log.Fatal(err)
	}
	return cardId
}

func EditCard(cardId primitive.ObjectID, cardUpdate CardUpdate) Card {
	updatedFields := bson.D{}
	if len(cardUpdate.Name) != 0 {
		updatedFields = append(updatedFields, bson.E{Key: "name", Value: cardUpdate.Name})
	}
	if len(cardUpdate.Description) != 0 {
		updatedFields = append(updatedFields, bson.E{Key: "description", Value: cardUpdate.Description})
	}
	if len(cardUpdate.AssignedTo) != 0 {
		updatedFields = append(updatedFields, bson.E{Key: "assignedTo", Value: cardUpdate.AssignedTo})
	}
	if len(cardUpdate.DueDate) != 0 {
		newDueDate, err := time.Parse("02-01-2006", cardUpdate.DueDate)
		if err != nil {
			log.Fatal(err)
		}
		updatedFields = append(updatedFields, bson.E{Key: "dueDate", Value: primitive.NewDateTimeFromTime(newDueDate)})
	}
	if len(updatedFields) == 0 {
		return GetCardById(cardId)
	}
	_, err := cardsCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": cardId},
		bson.D{
			{"$set", updatedFields},
			//{"$set", bson.D{{"name", "Updated but in code"}}},
		})
	if err != nil {
		log.Fatal(err)
	}

	return GetCardById(cardId)
}

func DeleteCard(cardId primitive.ObjectID) error {
	_, err := cardsCollection.DeleteOne(context.Background(), bson.D{{"_id", cardId}})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
