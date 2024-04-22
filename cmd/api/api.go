package api

import (
	"log"
	"net/http"

	"github.com/FrancoRutigliano/Mongo-go/middleware"
	"github.com/FrancoRutigliano/Mongo-go/services/todos"
)

type APIServer struct {
	addr string
}

// Constructor
func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	// enrutador principal
	router := http.NewServeMux()

	// enrutador secundario + /api/v1/
	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := middleware.MiddlewareChain()

	// todos
	todos.RegisterRoutes(router)

	log.Println("Listening on port: ", s.addr)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(subrouter),
	}

	return server.ListenAndServe()
}
