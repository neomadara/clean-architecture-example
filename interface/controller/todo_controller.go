package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neomadara/clean-architecture-example/usecase/interactor"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoController struct {
	todoInteractor interactor.TodoInteractor
}

func (tc todoController) FindTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, errObjectIdFromHex := primitive.ObjectIDFromHex(id)

	if errObjectIdFromHex != nil {
		return c.Status(404).JSON(errObjectIdFromHex)
	}

	todo, errUseCase := tc.todoInteractor.FindTodoById(objectId)
	if errUseCase != nil {
		return c.Status(500).JSON(errUseCase)
	}

	return c.Status(200).JSON(todo)
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
	FindTodoById(c *fiber.Ctx) error
}

func NewTodoController(ti interactor.TodoInteractor) TodoController {
	return &todoController{ti}
}
