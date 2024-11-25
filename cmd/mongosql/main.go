package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	todo "task-manager"
	hanMongo "task-manager/pkg/handler/mongo"
	repMongo "task-manager/pkg/repository/mongo"
	serMongo "task-manager/pkg/service/mongo"
	storMongo "task-manager/pkg/storage/mongo"
)

const (
	databaseName   = "data"
	collectionName = "languages"
)

type lang struct {
	ID   int
	Name string
}

func main() {
	db, err := storMongo.NewMongoDB(storMongo.Config{
		Host: "localhost",
		Port: "27017",
	})

	if err != nil {
		log.Fatalf("failed to connect to mongo db: %s", err.Error())
	}
	defer db.Disconnect(context.Background())

	// Проверка соединения
	if err != nil {
		fmt.Println("Не удалось подключиться к БД mongo:", err)
	} else {
		fmt.Println("Успешное подключение к БД mongo!")
	}

	err = db.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	repos := repMongo.NewRepository(db)
	services := serMongo.NewService(repos)
	handlers := hanMongo.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
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

	if err := db.Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
