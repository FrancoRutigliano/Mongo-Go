package todos

import (
	"context"
	"log"

	"github.com/FrancoRutigliano/Mongo-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	client *mongo.Client
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		client: client,
	}
}

func (s *Store) returnCollectionPointer(collection string) *mongo.Collection {
	return s.client.Database("todos_db").Collection(collection)
}

func (s *Store) GetAllTodos() ([]types.Todo, error) {
	collection := s.returnCollectionPointer("todos")

	var todos []types.Todo

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo types.Todo
		cursor.Decode(todo)
		todos = append(todos, todo)
	}

	return todos, nil
}
