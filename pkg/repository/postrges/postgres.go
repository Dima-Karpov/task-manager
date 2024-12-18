package postrges

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
)

const (
	usersTable      = "users"
	postsListsTable = "posts_lists"
	usersListsTable = "users_lists"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.Password,
			cfg.DBName,
			cfg.SSLMode,
		))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
