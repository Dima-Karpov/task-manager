package entities

import (
	"errors"
	"github.com/gofrs/uuid"
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

type UsersListMaria struct {
	Id     uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserId uuid.UUID `gorm:"type:uuid;not null"`
	PostId uuid.UUID `gorm:"type:uuid;not null"`
}

type PostListMaria struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserPostId uuid.UUID `gorm:"type:uuid;not null;index"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
}

type UpdatePostInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type PostResponseMaria struct {
	Id        uuid.UUID `json:"Id"`
	Title     string    `json:"Title"`
	Content   string    `json:"Content"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
