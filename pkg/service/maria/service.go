package service

import (
	"github.com/gofrs/uuid"
	"task-manager/internal/entities"
	"task-manager/pkg/repository/maria"
)

type Authorization interface {
	CreateUser(user entities.UserMaria) (uuid.UUID, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type PostList interface {
	Create(userId uuid.UUID, list entities.PostListMaria) (uuid.UUID, error)
	GetAll(userId uuid.UUID) ([]entities.PostListMaria, error)
	Delete(userId, id uuid.UUID) error
	GeById(userId, id uuid.UUID) (entities.PostResponseMaria, error)
	UpdatePost(userId, id uuid.UUID, input entities.UpdatePostInput) error
}

type Service struct {
	Authorization
	PostList
}

func NewService(repos *maria.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostList:      NewPostListService(repos.PostList),
	}
}
