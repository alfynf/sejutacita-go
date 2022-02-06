package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Client

func InitDB() {

	var err error
	config := os.Getenv("CONNECTION_STRING")

	clientOptions := options.Client().ApplyURI(config)

	DB, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
}
