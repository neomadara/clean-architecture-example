package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Desc      string             `json:"desc"`
	Completed bool               `json:"completed" default:"false"`
}
