package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neomadara/clean-architecture-example/usecase/interactor"
)

type todoController struct {
	todoInteractor interactor.TodoInteractor
}

func (tc todoController) GetAllTodos(c *fiber.Ctx) error {
	todos, err := tc.todoInteractor.GetAllTodos()

	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(todos)
}

type TodoController interface {
	GetAllTodos(c *fiber.Ctx) error
}

func NewTodoController(ti interactor.TodoInteractor) TodoController {
	return &todoController{ti}
}
