package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	todo "task-manager"
	"task-manager/internal/entities"
	handler "task-manager/pkg/handler/maria"
	repository "task-manager/pkg/repository/maria"
	service "task-manager/pkg/service/maria"
	"task-manager/pkg/storage/maria"
)

// Конфигурация
type Config struct {
	DSN string
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("error loading .env file")
	}

	cfg := Config{DSN: os.Getenv(viper.GetString("db.mariadsn"))}

	db, err := maria.NewMariaDB(maria.Config(cfg))
	if err != nil {
		log.Fatalln(err)
	}

	// Накатываем миграции
	err = db.AutoMigrate(&entities.UserMaria{}, entities.UsersListMaria{}, &entities.PostListMaria{})
	if err != nil {
		log.Fatalln("Failed to migrate database: ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if db != nil {
		sqlDB, err := db.DB() // Получаем низкоуровневое подключение sql.DB
		if err != nil {
			logrus.Fatalf("Error occured while connecting to database: %s", err.Error())
		} else {
			if err := sqlDB.Close(); err != nil {
				logrus.Fatalf("Error occured while closing database connection: %s", err.Error())
			} else {
				logrus.Print("TodoApp Stopped")
			}
		}
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
