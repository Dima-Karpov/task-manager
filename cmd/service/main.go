package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	postgres "task-manager/pkg/storage/postgres"

	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3333")

	// Чтение значений переменных окружения
	pwd := os.Getenv("POSTGRES_PASSWORD")
	user := os.Getenv("POSTGRES_USER")

	db, err := postgres.New(
		postgres.Config{
			Host:     "db",
			Port:     "5432",
			Username: user,
			Password: pwd,
			DBName:   "postgres",
			SSLMode:  "disable",
		},
	)

	if err != nil {
		log.Fatalf("failed to connect to postgres db: %s", err.Error())
	}

	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		fmt.Println("Не удалось подключиться к БД:", err)
	} else {
		fmt.Println("Успешное подключение к БД!")
	}
}
