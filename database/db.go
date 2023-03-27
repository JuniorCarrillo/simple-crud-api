package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type DB struct {
	client *mongo.Client
}

func Db() *DB {
	dbPassword := os.Getenv("MONGO_PASSWORD")
	dbUser := os.Getenv("MONGO_USERNAME")
	dbHost := os.Getenv("MONGO_HOST")
	dbPort := os.Getenv("MONGO_PORT")
	dbUri := os.Getenv("MONGO_URI")

	var uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/?retryWrites=true&w=majority", dbUser, dbPassword, dbHost, dbPort)
	if dbUri != "" {
		uri = dbUri
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{client: client}
}
