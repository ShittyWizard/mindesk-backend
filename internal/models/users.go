package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `bson:"id,omitempty"`
	FirstName string
	LastName  string
	Password  string
	Email     string
}

//todo: implement CRUD