package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neomadara/clean-architecture-example/interface/controller"
)

func NewRouter(app *fiber.App, c controller.AppController) *fiber.App {
	app.Get("/todos", func(ctx *fiber.Ctx) error { return c.Todo.GetAllTodos(ctx) })
	app.Get("/todos/:id", func(ctx *fiber.Ctx) error { return c.Todo.FindTodoById(ctx) })
	app.Post("/todos", func(ctx *fiber.Ctx) error { return c.Todo.CreateTodo(ctx) })
	app.Put("/todos/:id", func(ctx *fiber.Ctx) error { return c.Todo.UpdateTodo(ctx) })
	app.Delete("/todos/:id", func(ctx *fiber.Ctx) error { return c.Todo.DeleteTodo(ctx) })
	return app
}
