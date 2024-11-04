package memdb

import postgres "task-manager/pkg/storage/postgres"

// Пользовательский тип данных - реализация БД в памяти.
// Т.н. "заглушка".
type DB []postgres.Task

// Выполнение контракта интерфейса storage.Interface
func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}
func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}
