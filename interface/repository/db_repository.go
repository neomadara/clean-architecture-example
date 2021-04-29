package repository

import (
	"github.com/neomadara/clean-architecture-example/usecase/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbRepository struct {
	mongoDB *mongo.Database
}

func (d dbRepository) Transaction(f func(interface{}) (interface{}, error)) (interface{}, error) {
	panic("implement me")
}

func NewDBRepository(mongoDB *mongo.Database) repository.DBRepository {
	return &dbRepository{mongoDB}
}
