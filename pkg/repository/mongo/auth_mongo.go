package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task-manager/internal/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	usersTable = "users"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user entities.User) (string, error) {
	usersCollection := r.db.Database(usersTable).Collection(usersTable)

	result, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to cast InsertedID to ObjectID")
	}

	return objectID.Hex(), nil
}

func (r *AuthMongo) GetUser(username, password string) (entities.User, error) {
	var user entities.User
	filter := bson.D{{"username", username}, {"password", password}}
	usersCollection := r.db.Database(usersTable).Collection(usersTable)
	err := usersCollection.FindOne(context.Background(), filter).Decode(&user)

	return user, err
}
