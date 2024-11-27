package storage

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestTasks(t *testing.T) {
	// Создание фейковой базы данных для тестирования
	dbPool, err := pgxpool.Connect(context.Background(), "postgresql://task-manager:OvoIpFrIL2VS@localhost:5436/postgres")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("К базе подключился")
	defer dbPool.Close()

	storage := &Storage{db: dbPool}

	// Тестирование метода Tasks
	tasks, err := storage.Tasks(0, 0)
	assert.NoError(t, err)
	assert.NotNil(t, tasks)
}

func TestNewTask(t *testing.T) {
	// Создание фейковой базы данных для тестирования
	dbPool, err := pgxpool.Connect(context.Background(), "postgresql://task-manager:OvoIpFrIL2VS@localhost:5436/postgres")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("К базе подключился")
	defer dbPool.Close()

	storage := &Storage{db: dbPool}

	// Тестирование метода NewTask
	newTask := Task{
		Title:   "Test Task",
		Content: "This is a test task.",
	}
	id, err := storage.NewTask(newTask)
	assert.NoError(t, err)
	assert.Greater(t, id, 0)
}
