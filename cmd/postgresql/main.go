package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	todo "task-manager"
	"task-manager/pkg/handler/postgres"
	"task-manager/pkg/repository/postrges"
	"task-manager/pkg/service/postgres"
	postgres "task-manager/pkg/storage/postgres"

	"github.com/joho/godotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPostgresDB(
		postgres.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: os.Getenv(viper.GetString("db.username")),
			Password: os.Getenv(viper.GetString("db.password")),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	)

	if err != nil {
		log.Fatalf("failed to connect to postgres db: %s", err.Error())
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		fmt.Println("Не удалось подключиться к БД postgres:", err)
	} else {
		fmt.Println("Успешное подключение к БД postgres!")
	}

	repos := postrges.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http sever: %s", err.Error())
		}
	}()

	logrus.Print("TaskManagerApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TaskManagerApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shugging down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
