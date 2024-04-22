package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {
	//mongo collection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//getting username and password from the .env
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	// set auth credentials
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	//connect to mongo
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("conected to mongo")
	return client, nil
}
