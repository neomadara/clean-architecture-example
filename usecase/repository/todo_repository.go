package repository

import "github.com/neomadara/clean-architecture-example/domain/model"

type TodoRepository interface {
	GetAllTodos() ([]*model.Todo, error)
	FindTodoById(id string) (*model.Todo, error)
}
