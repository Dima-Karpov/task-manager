package service

import (
	"task-manager/internal/entities"
	"task-manager/pkg/repository/mongo"
)

type Authorization interface {
	CreateUser(user entities.UserMongo) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type PostList interface {
	Create(userId string, post entities.PostListMongo) (string, error)
	GetAll(userId string) ([]entities.PostListMongo, error)
	GetById(userId, id string) (entities.PostListMongo, error)
	Delete(userId, id string) error
	Update(userId, id string, post entities.UpdatePostInput) error
}

type Service struct {
	Authorization
	PostList
}

func NewService(repos *mongo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostList:      NewPostListService(repos.PostList),
	}
}
