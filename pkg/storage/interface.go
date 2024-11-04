package storage

import postgres "task-manager/pkg/storage/postgres"

// Интерфейс БД
// Этот интерфейс позволяет абстрагироваться конкретной СУБД.
// Можно создать реализацию БД в памяти для модульных тестов.

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
}
