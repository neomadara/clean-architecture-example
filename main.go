package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neomadara/clean-architecture-example/infrastructure/datastore"
	"github.com/neomadara/clean-architecture-example/infrastructure/router"
	"github.com/neomadara/clean-architecture-example/registry"
)

func main() {
	mongoClient := datastore.MongoClient()
	r := registry.NewRegistry(mongoClient)

	app := fiber.New()
	app = router.NewRouter(app, r.NewAppController())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
