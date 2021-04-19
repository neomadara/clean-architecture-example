package repository

import (
	"context"
	"github.com/neomadara/clean-architecture-example/domain/model"
	"github.com/neomadara/clean-architecture-example/usecase/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type todoRepository struct {
	mongoClient *mongo.Client
}

func (db todoRepository) GetAllTodos() ([]*model.Todo, error) {
	var ctx = context.TODO()
	var todos []*model.Todo

	database := db.mongoClient.Database("todosapp")
	collection := database.Collection("todos")

	filter := bson.D{}
	result, errDB := collection.Find(ctx, filter)

	if errDB != nil {
		return nil, errDB
	}

	for result.Next(ctx) {
		var todo *model.Todo
		err := result.Decode(&todo)
		if err != nil {
			log.Print(err)
		}
		todos = append(todos, todo)
	}

	return todos, nil

}

func NewTodoRepository(mongoClient *mongo.Client) repository.TodoRepository {
	return &todoRepository{mongoClient}
}
