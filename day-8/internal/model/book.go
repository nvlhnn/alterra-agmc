package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
