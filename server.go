package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/brandon-julio-t/graph-gongular-backend/graph"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/generated"
)

const defaultPort = "8080"
const graphqlEndpoint = "/graphql"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := setupServer()
	setupServerRoutings(srv)
	runServer(port)
}

func setupServer() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
}

func setupServerRoutings(srv *handler.Server) {
	http.Handle("/", playground.Handler("GraphQL playground", graphqlEndpoint))
	http.Handle(graphqlEndpoint, srv)
}

func runServer(port string) {
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
