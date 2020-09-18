package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Card struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	AssignedTo  string             `bson:"assignedTo,omitempty"`
	DueDate     string             `bson:"dueDate,omitempty"`
}

type CardUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AssignedTo  string `bson:"assignedTo,omitempty"`
	DueDate     string `bson:"dueDate,omitempty"`
}

var cardsCollectionName = "cards"
var cardsCollection *mongo.Collection

func InitCardsCollection() {
	log.Println("Trying to connect to MongoDB...")
	cardsCollection = mongodb.Collection(cardsCollectionName)
	log.Println("Connected to MongoDB...")
}

// If you change size if test cards - test it too in test/models/cards_test.go:17
func AddTestCards() {
	cardsCollection.DeleteMany(context.Background(), bson.D{})
	testCard1 := Card{primitive.NewObjectID(), "test name 1", "test description 1", "me1", "10.08.2021"}
	AddCard(testCard1)

	testCard2 := Card{primitive.NewObjectID(), "test name 2", "test description 2", "me2", "11.08.2021"}
	AddCard(testCard2)

	testCard3 := Card{primitive.NewObjectID(), "test name 3", "test description 3", "me3", "12.08.2021"}
	AddCard(testCard3)

	testCard4 := Card{primitive.NewObjectID(), "test name 4", "test description 4", "me4", "13.08.2021"}
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

func AddCard(card Card) primitive.ObjectID {
	card.Id = primitive.NewObjectID()
	_, err := cardsCollection.InsertOne(context.Background(), card)
	if err != nil {
		log.Fatal(err)
	}
	return card.Id
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
		updatedFields = append(updatedFields, bson.E{Key: "dueDate", Value: cardUpdate.DueDate})
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
