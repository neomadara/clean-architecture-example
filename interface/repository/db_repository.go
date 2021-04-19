package repository

import (
	"github.com/neomadara/clean-architecture-example/usecase/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type dbRepository struct {
	mongoClient *mongo.Client
}

func (d dbRepository) Transaction(f func(interface{}) (interface{}, error)) (interface{}, error) {
	panic("implement me")
}

func NewDBRepository(mongoClient *mongo.Client) repository.DBRepository {
	return &dbRepository{mongoClient}
}
