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

func InitMongoDB(dataSource string) {
	clientOptions := options.Client().ApplyURI(dataSource)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	mongodb = client.Database("mindesk")
}
