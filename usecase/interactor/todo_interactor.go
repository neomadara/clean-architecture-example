package interactor

import (
	"github.com/neomadara/clean-architecture-example/domain/model"
	"github.com/neomadara/clean-architecture-example/usecase/presenter"
	"github.com/neomadara/clean-architecture-example/usecase/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoInteractor struct {
	TodoRepository repository.TodoRepository
	TodoPresenter  presenter.TodoPresenter
	DBRepository   repository.DBRepository
}

func (us *todoInteractor) DeleteTodo(id primitive.ObjectID) error {
	return us.TodoRepository.DeleteTodo(id)
}

func (us *todoInteractor) UpdateTodo(id primitive.ObjectID, todo *model.Todo) (*model.Todo, error) {
	todo, err := us.TodoRepository.UpdateTodo(id, todo)
	if err != nil {
		return nil, err
	}
	return us.TodoPresenter.ResponseTodo(todo), nil
}

func (us *todoInteractor) CreateTodo(todo *model.Todo) error {
	return us.TodoRepository.CreateTodo(todo)
}

func (us *todoInteractor) FindTodoById(id primitive.ObjectID) (*model.Todo, error) {
	todo, err := us.TodoRepository.FindTodoById(id)
	if err != nil {
		return nil, err
	}
	return us.TodoPresenter.ResponseTodo(todo), nil
}

func (us *todoInteractor) GetAllTodos() ([]*model.Todo, error) {
	todos, err := us.TodoRepository.GetAllTodos()
	if err != nil {
		return nil, err
	}
	return us.TodoPresenter.ResponseTodos(todos), nil
}

type TodoInteractor interface {
	GetAllTodos() ([]*model.Todo, error)
	FindTodoById(id primitive.ObjectID) (*model.Todo, error)
	CreateTodo(*model.Todo) error
	UpdateTodo(id primitive.ObjectID, todo *model.Todo) (*model.Todo, error)
	DeleteTodo(id primitive.ObjectID) error
}

func NewTodoInteractor(r repository.TodoRepository, p presenter.TodoPresenter, d repository.DBRepository) TodoInteractor {
	return &todoInteractor{r, p, d}
}
