package server

import (
	"fmt"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	DB  *mongo.Client
    MUX *http.ServeMux
}

func (s *Server) Run() {
    s.MUX = http.NewServeMux()
    RoutesHandlers(s.MUX)

	httpServer := &http.Server{
		Addr:    ":3001",
		Handler: s.MUX,
	}

	fmt.Printf("Server listening on http port 3001")
	log.Fatal(httpServer.ListenAndServe())
}
