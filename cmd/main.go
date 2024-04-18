package main

import (
	"log"

	"github.com/FrancoRutigliano/Mongo-go/cmd/api"
)

func main() {
	PORT := ":8080"
	server := api.NewApiServer(PORT)
	if err := server.Run(); err != nil {
		log.Fatal("error to initialize the server on the port: ", PORT)
	}
}
