package entities

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostList struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
type CreatePostListBody struct {
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type UpdatePostListBody struct {
	Title   string
	Content string
}

type PostListMongo struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UsersList struct {
	Id     int
	UserId int
	PostId int
}

type UpdatePostInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
