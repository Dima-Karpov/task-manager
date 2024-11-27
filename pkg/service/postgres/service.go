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

type Service struct {
	Authorization
}

func NewService(repos *postrges.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
