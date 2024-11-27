package service

import (
	"task-manager/internal/entities"
	"task-manager/pkg/repository/mongo"
)

type Authorization interface {
	CreateUser(user entities.User) (string, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *mongo.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
