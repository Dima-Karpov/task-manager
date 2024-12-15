package entities

import (
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserMongo struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" binding:"required"`
	Username string             `json:"username" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

type UserMaria struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Username string    `gorm:"type:varchar(255);not null"`
	Password string    `gorm:"type:varchar(1000);not null"`
}
