package handler

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task-manager/internal/entities"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"task-manager/pkg/service/mongo"
)

const salt = "hjqrhjqw124617ajfhajs"

// MockAuthorizationRepo — мок для репозитория mongo.Authorization
type MockAuthorizationRepo struct {
	mock.Mock
}

func (m *MockAuthorizationRepo) CreateUser(user entities.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}
func (m *MockAuthorizationRepo) GetUser(username, password string) (entities.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(entities.User), args.Error(1)
}

// generatePasswordHash — вспомогательная функция для хэширования
func generatePasswordHash(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func TestSignUp(t *testing.T) {
	// Настройка моков
	mockRepo := new(MockAuthorizationRepo)

	// Создаем AuthService с моком репозитория
	authService := service.NewAuthService(mockRepo)

	// Создаем Handler с AuthService
	handler := &Handler{
		services: &service.Service{
			Authorization: authService,
		},
	}

	// Настраиваем маршрутизатор
	router := gin.Default()
	router.POST("/signup", handler.signUp)

	// Тестовые данные
	testUser := entities.User{
		Name:     "Test Name",
		Username: "test_username",
		Password: "test_password",
	}

	hashedPassword := generatePasswordHash(testUser.Password, salt) // Считаем ожидаемый хэш

	expectedUser := entities.User{
		Name:     testUser.Name,
		Username: testUser.Username,
		Password: hashedPassword, // Используем хэш вместо исходного пароля
	}

	// Мокаем поведение репозитория
	mockRepo.On("CreateUser", expectedUser).Return("user_id_123", nil)

	// Сериализуем тестовые данные в JSON
	jsonData, err := json.Marshal(testUser)
	assert.NoError(t, err)

	// Выполняем запрос
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверяем результат
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "user_id_123", response["id"]) // Проверяем, что ID корректен
	mockRepo.AssertExpectations(t)                 // Убедиться, что мок был вызван с правильными параметрами
}
