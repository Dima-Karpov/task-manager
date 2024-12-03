package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"task-manager/internal/entities"
	customErrors "task-manager/pkg/errors"
	"time"
)

type PostListMongo struct {
	db *mongo.Client
}

func NewPostListMongo(db *mongo.Client) *PostListMongo {
	return &PostListMongo{db: db}
}

func (m *PostListMongo) Create(userId string, post entities.PostListMongo) (string, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", errors.New("invalid user id")
	}

	var user entities.UserMongo
	usersCollection := m.db.Database(usersTable).Collection(usersTable)
	filter := bson.D{{"_id", userObjectId}}
	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("user not found: " + userId)
		}
		return "", errors.New("failed to find user")
	}

	// Add new post
	post.Id = primitive.NewObjectID()
	post.UserId = userObjectId
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	postsCollection := m.db.Database(postsTable).Collection(postsTable)
	_, err = postsCollection.InsertOne(context.Background(), post)
	if err != nil {
		return "", errors.New("failed to insert post: " + err.Error())
	}

	return post.Id.Hex(), nil
}

func (m *PostListMongo) GetAll(userId string) ([]entities.PostListMongo, error) {
	var lists []entities.PostListMongo
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return lists, errors.New("invalid user id")
	}

	var user entities.UserMongo
	usersCollection := m.db.Database(usersTable).Collection(usersTable)
	filter := bson.D{{"_id", userObjectId}}

	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return lists, errors.New("user not found: " + userId)
		}
		return lists, errors.New("failed to find user: " + err.Error())
	}

	postsCollection := m.db.Database(postsTable).Collection(postsTable)
	postsFilter := bson.M{"userId": userObjectId}

	cursor, err := postsCollection.Find(context.Background(), postsFilter)
	if err != nil {
		return lists, errors.New("failed to find posts: " + err.Error())
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &lists); err != nil {
		return lists, errors.New("failed to decode posts: " + err.Error())
	}

	return lists, nil
}

func (m *PostListMongo) GetById(userId, id string) (entities.PostListMongo, error) {
	var list entities.PostListMongo
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return list, errors.New("invalid user id")
	}
	postObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return list, errors.New("invalid post id")
	}

	postsCollection := m.db.Database(postsTable).Collection(postsTable)
	filter := bson.M{"_id": postObjectId, "userId": userObjectId}

	err = postsCollection.FindOne(context.Background(), filter).Decode(&list)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("Post not found: %v", filter)
			return list, errors.New("post not found: " + id)
		}
		return list, errors.New("failed to find post: " + err.Error())
	}

	return list, nil
}

func (m *PostListMongo) Delete(userId, id string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return errors.New("invalid user id")
	}
	postObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid post id")
	}

	postsCollection := m.db.Database(postsTable).Collection(postsTable)
	filter := bson.M{
		"_id":    postObjectId,
		"userId": userObjectId,
	}

	result, err := postsCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return errors.New("failed to delete post: " + err.Error())
	}

	if result.DeletedCount == 0 {
		return customErrors.ErrPostNotFound
	}

	return nil
}

func (m *PostListMongo) Update(userId, id string, post entities.UpdatePostInput) error {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return errors.New("invalid user id")
	}
	postObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid post id")
	}

	var user entities.UserMongo
	usersCollection := m.db.Database(usersTable).Collection(usersTable)
	filter := bson.D{{"_id", userObjectId}}

	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("user not found: " + userId)
		}
		return errors.New("failed to find user: " + err.Error())
	}

	updateData := bson.M{}
	if post.Title != nil {
		updateData["title"] = *post.Title
	}
	if post.Content != nil {
		updateData["content"] = *post.Content
	}
	updateData["updated_at"] = time.Now()

	postsCollection := m.db.Database(postsTable).Collection(postsTable)
	postFilter := bson.M{
		"_id":    postObjectId,
		"userId": userObjectId,
	}

	update := bson.M{"$set": updateData}

	result, err := postsCollection.UpdateOne(context.Background(), postFilter, update)
	if err != nil {
		return errors.New("post not found or does not belong to user")
	}
	if result.MatchedCount == 0 {
		return customErrors.ErrPostNotFound
	}

	return nil
}
