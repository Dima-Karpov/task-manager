package postrges

import (
	"github.com/jmoiron/sqlx"
	"task-manager/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
}

type PostList interface {
	Create(userId int, post entities.PostList) (int, error)
	GetAll(userId int) ([]entities.PostList, error)
	GetById(userId, id int) (entities.PostList, error)
	Delete(userId, id int) error
	Update(userId, id int, input entities.UpdatePostInput) error
}

type Repository struct {
	Authorization
	PostList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PostList:      NewPostListPostgres(db),
	}
}
