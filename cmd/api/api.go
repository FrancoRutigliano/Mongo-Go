package api

import (
	"log"
	"net/http"

	"github.com/FrancoRutigliano/Mongo-go/middleware"
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
	router.HandleFunc("GET /Todos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All todos"))
	})

	log.Println("Listening on port: ", s.addr)

	server := &http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(subrouter),
	}

	return server.ListenAndServe()
}
