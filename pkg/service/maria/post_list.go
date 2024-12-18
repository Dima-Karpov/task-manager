package service

import (
	"github.com/gofrs/uuid"
	"task-manager/internal/entities"
	repository "task-manager/pkg/repository/maria"
)

type PostListService struct {
	repo repository.PostList
}

func NewPostListService(repo repository.PostList) *PostListService {
	return &PostListService{repo: repo}
}
func (s *PostListService) Create(userId uuid.UUID, list entities.PostListMaria) (uuid.UUID, error) {
	return s.repo.Create(userId, list)
}
func (s *PostListService) GetAll(userId uuid.UUID) ([]entities.PostListMaria, error) {
	return s.repo.GetAll(userId)
}
func (s *PostListService) Delete(userId, id uuid.UUID) error {
	return s.repo.Delete(userId, id)
}
func (s *PostListService) GeById(userId, id uuid.UUID) (entities.PostResponseMaria, error) {
	return s.repo.GeById(userId, id)
}
func (s *PostListService) UpdatePost(userId, id uuid.UUID, input entities.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdatePost(userId, id, input)
}
