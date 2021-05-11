package repository

import (
	"github.com/neomadara/clean-architecture-example/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	GetAllTodos() ([]*model.Todo, error)
	FindTodoById(id primitive.ObjectID) (*model.Todo, error)
	CreateTodo(*model.Todo) error
	UpdateTodo(id primitive.ObjectID, todo *model.Todo) (*model.Todo, error)
}
