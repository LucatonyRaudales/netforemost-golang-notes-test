package server

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	DB  *mongo.Client
    MUX *http.ServeMux
}
func (s *Server) Run() {
	var err error
	s.DB, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://hindra:vUxNd2AF3pTWHkgb@hindra-auth.b8lhl8z.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	err = s.DB.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer s.DB.Disconnect(context.Background())
    s.MUX = http.NewServeMux()
    s.RoutesHandlers(s.MUX)

	httpServer := &http.Server{
		Addr:    ":3001",
		Handler: s.MUX,
	}

	fmt.Printf("Server listening on http port 3001")
	log.Fatal(httpServer.ListenAndServe())
}
