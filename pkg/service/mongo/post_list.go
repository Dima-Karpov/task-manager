package service

import (
	"task-manager/internal/entities"
	repository "task-manager/pkg/repository/mongo"
)

type PostListService struct {
	repo repository.PostList
}

func NewPostListService(repo repository.PostList) *PostListService {
	return &PostListService{repo: repo}
}
func (s *PostListService) Create(userId string, post entities.PostListMongo) (string, error) {
	return s.repo.Create(userId, post)
}
func (s *PostListService) GetAll(userId string) ([]entities.PostListMongo, error) {
	return s.repo.GetAll(userId)
}
func (s *PostListService) GetById(userId, id string) (entities.PostListMongo, error) {
	return s.repo.GetById(userId, id)
}
func (s *PostListService) Delete(userId, id string) error {
	return s.repo.Delete(userId, id)
}
func (s *PostListService) Update(userId, id string, post entities.UpdatePostInput) error {
	return s.repo.Update(userId, id, post)
}
