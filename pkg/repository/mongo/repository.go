package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"task-manager/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (string, error)
	GetUser(username, password string) (entities.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
