package service

import (
	"task-manager/internal/entities"
	repository "task-manager/pkg/repository/postrges"
)

type PostListService struct {
	repo repository.PostList
}

func NewPostListService(repo repository.PostList) *PostListService {
	return &PostListService{repo: repo}
}
func (s *PostListService) Create(userId int, task entities.PostList) (int, error) {
	return s.repo.Create(userId, task)
}
func (s *PostListService) GetAll(userId int) ([]entities.PostList, error) {
	return s.repo.GetAll(userId)
}
func (s *PostListService) GetById(userId, id int) (entities.PostList, error) {
	return s.repo.GetById(userId, id)
}
func (s *PostListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
func (s *PostListService) Update(userId, id int, input entities.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, id, input)
}
