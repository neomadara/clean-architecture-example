package repository

import (
	"context"
	"github.com/neomadara/clean-architecture-example/domain/model"
	"github.com/neomadara/clean-architecture-example/usecase/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type todoRepository struct {
	mongoClient *mongo.Client
}

func (db todoRepository) CreateTodo(todo *model.Todo) error {
	var ctx = context.TODO()

	_, err := db.mongoClient.Database("todosapp").Collection("todos").InsertOne(ctx, todo)
	if err != nil {
		return err
	}
	return nil
}

func (db todoRepository) FindTodoById(id primitive.ObjectID) (*model.Todo, error) {
	var ctx = context.TODO()
	var todo *model.Todo

	filter := bson.D{{"_id", id}}
	database := db.mongoClient.Database("todosapp")
	collection := database.Collection("todos")

	result := collection.FindOne(ctx, filter)

	err := result.Decode(&todo)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return todo, nil
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
