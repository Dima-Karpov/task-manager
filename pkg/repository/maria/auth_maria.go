package maria

import (
	"errors"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"task-manager/internal/entities"
)

type AuthMaria struct {
	db *gorm.DB
}

func NewAuthMaria(db *gorm.DB) *AuthMaria {
	return &AuthMaria{db: db}
}

func (m *AuthMaria) CreateUser(user entities.UserMaria) (uuid.UUID, error) {
	var existingUser entities.UserMaria
	err := m.db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		// Пользователь найден, возвращаем ошибку
		return uuid.Nil, errors.New("username already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Если ошибка не "record not found", возвращаем её
		return uuid.Nil, err
	}

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err // Вернуть ошибку, если генерация UUID не удалась
	}
	user.Id = newUUID

	if err := m.db.Create(&user).Error; err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}

func (m *AuthMaria) GetUser(username, password string) (entities.UserMaria, error) {
	var user entities.UserMaria
	if err := m.db.Where(
		"username = ? AND password = ?",
		username,
		password,
	).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found or invalid credentials")
		}
		return user, err
	}

	return user, nil
}
