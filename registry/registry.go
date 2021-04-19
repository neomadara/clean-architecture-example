package registry

import (
	"github.com/neomadara/clean-architecture-example/interface/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type registry struct {
	mongoClient *mongo.Client
}

type Registry interface {
	NewAppController() controller.AppController
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Todo: r.NewTodoController(),
	}
}

func NewRegistry(mongoClient *mongo.Client) Registry {
	return &registry{mongoClient}
}
