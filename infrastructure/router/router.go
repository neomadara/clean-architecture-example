package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neomadara/clean-architecture-example/interface/controller"
)

func NewRouter(app *fiber.App, c controller.AppController) *fiber.App {
	app.Get("/todos", func(ctx *fiber.Ctx) error { return c.Todo.GetAllTodos(ctx) })
	app.Get("/todos/:id", func(ctx *fiber.Ctx) error { return c.Todo.FindTodoById(ctx) })
	return app
}
