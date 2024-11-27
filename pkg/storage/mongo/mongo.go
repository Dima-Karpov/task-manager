package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Config struct {
	Host string
	Port string
}

func NewMongoDB(cfg Config) (*mongo.Client, error) {
	url := "mongodb://" + cfg.Host + ":" + cfg.Port

	mongoOpts := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), mongoOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
