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
	router.HandleFunc("GET /todos", handleGetTodos)
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola mundo"))
}
