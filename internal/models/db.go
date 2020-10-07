package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongodb *mongo.Database

// Mongo address for only docker (it could be run by scripts/start-docker-test-env.sh
var TestMongoAddress = "mongodb://localhost:27017"

// Mongo address in docker-compose
var MongoAddress = "mongodb://mongodb:27017"

func InitMongoDB(dataSource string) {
	clientOptions := options.Client().ApplyURI(dataSource)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	mongodb = client.Database("mindesk")

	// Init collections
	InitCardsCollection()
	InitColumnsCollection()
	InitDesksCollection()

	// Init test data
	testDeskId := AddTestDesks()[0]
	testColumnId := AddTestColumns(testDeskId)
	AddTestCards(testColumnId, testDeskId)
}
