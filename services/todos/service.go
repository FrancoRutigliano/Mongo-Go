package todos

import (
	"github.com/FrancoRutigliano/Mongo-go/types"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) types.Todo {
	client = mongo

	return types.Todo{}
}
