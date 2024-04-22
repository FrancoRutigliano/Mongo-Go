package main

import (
	"context"
	"log"
	"time"

	"github.com/FrancoRutigliano/Mongo-go/cmd/api"
	"github.com/FrancoRutigliano/Mongo-go/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	mongoClient, err := initStorage()
	if err != nil {
		log.Panic(err)
	}

	PORT := ":8080"
	server := api.NewApiServer(PORT, mongoClient)
	if err := server.Run(); err != nil {
		log.Fatal("error to initialize the server on the port: ", PORT)
	}
}

func initStorage() (*mongo.Client, error) {
	// conectarse a mongo
	client, err := db.ConnectToMongo()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("DB: Succesfully conected")

	return client, nil
}
