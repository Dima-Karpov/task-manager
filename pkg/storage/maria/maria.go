package maria

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	DSN string
}

func NewMariaDB(cfg Config) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB() // Получаем объект *sql.DB из *gorm.DB
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
