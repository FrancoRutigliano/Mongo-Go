package todos

import (
	"net/http"

	"github.com/FrancoRutigliano/Mongo-go/types"
)

type Handler struct {
	store types.TodoStore
}

// constructor
func NewHandler(store types.TodoStore) *Handler {
	return &Handler{
		store: store,
	}
}

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /todos", handleGetAllTodos)
	router.HandleFunc("POST /todos", handleCreateTodos)
	router.HandleFunc("PUT /todos/update/{id}", handleUpdateTodos)
	router.HandleFunc("DELETE /todos/delete/{id}", handleDeleteTodos)
}

func handleGetAllTodos(w http.ResponseWriter, r *http.Request) {

}

func handleCreateTodos(w http.ResponseWriter, r *http.Request) {
	// get json payload

	//validate the payload

	// call the db endpoint
}

func handleUpdateTodos(w http.ResponseWriter, r *http.Request) {
	// take the id of the params

	// get json payload

	// validate payload

	// call db endpoint
}

func handleDeleteTodos(w http.ResponseWriter, r *http.Request) {
	// take the id of the params
}
