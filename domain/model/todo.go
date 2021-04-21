package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Title     string             `json:"title"`
	Desc      string             `json:"description"`
	Completed bool               `json:"completed"`
}
