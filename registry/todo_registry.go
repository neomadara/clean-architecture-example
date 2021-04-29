package registry

import (
	"github.com/neomadara/clean-architecture-example/interface/controller"
	"github.com/neomadara/clean-architecture-example/usecase/interactor"

	ip "github.com/neomadara/clean-architecture-example/interface/presenter"
	ir "github.com/neomadara/clean-architecture-example/interface/repository"
	tp "github.com/neomadara/clean-architecture-example/usecase/presenter"
	tr "github.com/neomadara/clean-architecture-example/usecase/repository"
)

func (r *registry) NewTodoController() controller.TodoController {
	return controller.NewTodoController(r.NewTodoInteractor())
}

func (r *registry) NewTodoInteractor() interactor.TodoInteractor {
	return interactor.NewTodoInteractor(r.NewTodoRepository(), r.NewTodoPresenter(), ir.NewDBRepository(r.mongoDB))
}

func (r *registry) NewTodoRepository() tr.TodoRepository {
	return ir.NewTodoRepository(r.mongoDB)
}

func (r *registry) NewTodoPresenter() tp.TodoPresenter {
	return ip.NewTodoPresenter()
}
