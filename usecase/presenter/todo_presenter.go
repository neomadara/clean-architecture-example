package presenter

import "github.com/neomadara/clean-architecture-example/domain/model"

type TodoPresenter interface {
	ResponseTodos(r []*model.Todo) []*model.Todo
	ResponseTodo(r *model.Todo) *model.Todo
}
