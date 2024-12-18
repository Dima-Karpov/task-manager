package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"task-manager/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.UserMongo) (string, error)
	GetUser(username, password string) (entities.UserMongo, error)
}

type PostList interface {
	Create(userId string, post entities.PostListMongo) (string, error)
	GetAll(userId string) ([]entities.PostListMongo, error)
	GetById(userId, id string) (entities.PostListMongo, error)
	Delete(userId, id string) error
	Update(userId, id string, post entities.UpdatePostInput) error
}

type Repository struct {
	Authorization
	PostList
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		PostList:      NewPostListMongo(db),
	}
}
