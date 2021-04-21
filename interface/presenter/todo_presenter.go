package presenter

import (
	"github.com/neomadara/clean-architecture-example/domain/model"
	"github.com/neomadara/clean-architecture-example/usecase/presenter"
)

type todoPresenter struct{}

func (tp todoPresenter) ResponseTodo(r *model.Todo) *model.Todo {
	return r
}

func (tp todoPresenter) ResponseTodos(r []*model.Todo) []*model.Todo {
	return r
}

func NewTodoPresenter() presenter.TodoPresenter {
	return &todoPresenter{}
}
