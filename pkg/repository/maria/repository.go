package maria

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"task-manager/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.UserMaria) (uuid.UUID, error)
	GetUser(username, password string) (entities.UserMaria, error)
}
type PostList interface {
	Create(userId uuid.UUID, list entities.PostListMaria) (uuid.UUID, error)
	GetAll(userId uuid.UUID) ([]entities.PostListMaria, error)
	Delete(userId, id uuid.UUID) error
	GeById(userId, id uuid.UUID) (entities.PostResponseMaria, error)
	UpdatePost(userId, id uuid.UUID, input entities.UpdatePostInput) error
}

type Repository struct {
	Authorization
	PostList
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMaria(db),
		PostList:      NewPostListMaria(db),
	}
}
