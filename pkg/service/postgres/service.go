package service

import (
	"task-manager/internal/entities"
	"task-manager/pkg/repository/postrges"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type PostList interface {
	Create(userId int, post entities.PostList) (int, error)
	GetAll(userId int) ([]entities.PostList, error)
	GetById(userId, id int) (entities.PostList, error)
	Delete(userId, id int) error
	Update(userId, id int, input entities.UpdatePostInput) error
}

type Service struct {
	Authorization
	PostList
}

func NewService(repos *postrges.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostList:      NewPostListService(repos.PostList),
	}
}
