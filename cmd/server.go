package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ebalkanski/graphql/internal/clients"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/ebalkanski/graphql/graph"
	"github.com/ebalkanski/graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//create users client
	usersClient, err := clients.New("grpc:8081")
	if err != nil {
		log.Printf("error creating users client: %v", err)
	}
	defer func() {
		if err := usersClient.Close(); err != nil {
			log.Print("closing users client failed: %v", err)
		}
	}()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(usersClient)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
