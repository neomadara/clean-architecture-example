package presenter

import "github.com/neomadara/clean-architecture-example/domain/model"

type TodoPresenter interface {
	ResponseTodo(r []*model.Todo) []*model.Todo
}
