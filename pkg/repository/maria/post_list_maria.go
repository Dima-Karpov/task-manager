package maria

import (
	"errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"task-manager/internal/entities"
	"time"
)

type PostListMaria struct {
	db *gorm.DB
}

func NewPostListMaria(db *gorm.DB) *PostListMaria {
	return &PostListMaria{
		db: db,
	}
}

func (m *PostListMaria) Create(userId uuid.UUID, list entities.PostListMaria) (uuid.UUID, error) {
	postId := uuid.Must(uuid.NewV4())
	userPostId := uuid.Must(uuid.NewV4())
	list.Id = postId
	list.UserPostId = userPostId
	list.CreatedAt = time.Now()
	list.UpdatedAt = time.Now()

	if err := m.db.Create(&list).Error; err != nil {
		return uuid.Nil, errors.New("failed to create post: " + err.Error())
	}

	userPost := entities.UsersListMaria{
		Id:     userPostId,
		UserId: userId,
		PostId: postId,
	}

	if err := m.db.Create(&userPost).Error; err != nil {
		return uuid.Nil, errors.New("failed to create user-post mapping: " + err.Error())
	}

	return postId, nil
}

func (m *PostListMaria) GetAll(userId uuid.UUID) ([]entities.PostListMaria, error) {
	var posts []entities.PostListMaria
	var userPosts []entities.UsersListMaria
	err := m.db.Where("user_id = ?", userId).Find(&userPosts).Error
	if err != nil {
		return posts, err
	}

	if len(userPosts) == 0 {
		return posts, errors.New("no posts found for user " + userId.String())
	}

	postIds := make([]uuid.UUID, len(userPosts))
	for i, userPost := range userPosts {
		postIds[i] = userPost.PostId
	}

	err = m.db.Where("id IN (?)", postIds).Find(&posts).Error
	if err != nil {
		return posts, errors.New("failed to find all posts: " + err.Error())
	}

	return posts, nil
}

func (m *PostListMaria) Delete(userId, id uuid.UUID) error {
	var userPost entities.UsersListMaria
	err := m.db.Where("user_id = ? AND post_id = ?", userId, id).First(&userPost).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found for the user")
		}
		return err
	}

	err = m.db.Where("id = ?", id).Delete(&entities.PostListMaria{}).Error
	if err != nil {
		return errors.New("failed to delete post: " + err.Error())
	}

	err = m.db.Where("post_id = ?", id).Delete(&entities.UsersListMaria{}).Error
	if err != nil {
		return errors.New("failed to delete user-post mapping: " + err.Error())
	}

	return nil
}

func (m *PostListMaria) GeById(userId, id uuid.UUID) (entities.PostResponseMaria, error) {
	var response entities.PostResponseMaria
	var post entities.PostListMaria
	var userPost entities.UsersListMaria

	err := m.db.Where("user_id = ? AND post_id = ?", userId, id).First(&userPost).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, errors.New("post not found for the user")
		}
		return response, err
	}

	err = m.db.Where("id = ?", id).First(&post).Error
	if err != nil {
		return response, errors.New("failed to find post: " + err.Error())
	}

	response = entities.PostResponseMaria{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return response, nil
}
func (m *PostListMaria) UpdatePost(userId, id uuid.UUID, input entities.UpdatePostInput) error {
	var userPost entities.UsersListMaria
	var post entities.PostListMaria

	err := m.db.Where("user_id = ? AND post_id = ?", userId, id).First(&userPost).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found for the user")
		}
		return err
	}

	err = m.db.First(&post, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Content != nil {
		post.Content = *input.Content
	}
	post.UpdatedAt = time.Now()

	err = m.db.Save(&post).Error
	if err != nil {
		return errors.New("failed to update post: " + err.Error())
	}

	return nil
}
