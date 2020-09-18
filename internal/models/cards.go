package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Card struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AssignedTo string `json:"assignedTo,omitempty"`
	DueDate string `json:"dueDate,omitempty"`
}

var cardsCollectionName = "cards"
var cardsCollection *mongo.Collection

func InitCardsCollection() {
	log.Println("Trying to connect to MongoDB...")
	cardsCollection = mongodb.Collection(cardsCollectionName)
	log.Println("Connected to MongoDB...")
}

func AddTestCards() {
	cardsCollection.DeleteMany(context.Background(), bson.D{})
	testCard1 := Card{primitive.NewObjectID(), "test name 1", "test description 1", "me", "10.08.2021"}
	AddCard(testCard1)

	testCard2 := Card{primitive.NewObjectID(), "test name 2", "test description 2", "me", "11.08.2021"}
	AddCard(testCard2)

	testCard3 := Card{primitive.NewObjectID(), "test name 3", "test description 3", "me", "12.08.2021"}
	AddCard(testCard3)

	testCard4 := Card{primitive.NewObjectID(), "test name 4", "test description 4", "me", "13.08.2021"}
	AddCard(testCard4)
}

func AllCards() []*Card {
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

func AddCard(card Card) interface{} {
	insertedCard, err := cardsCollection.InsertOne(context.Background(), card)
	if err != nil {
		log.Fatal(err)
	}
	return insertedCard.InsertedID
}

// todo: implement it
//func EditCard() Card{
//	initCardsCollection()
//}

func DeleteCard(cardId string) error {
	_, err := cardsCollection.DeleteOne(context.Background(), bson.M{"_id": cardId})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
