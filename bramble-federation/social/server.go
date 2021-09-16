package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ihac/graphql-poc-playground/bramble-federation/social/graph"
	"github.com/ihac/graphql-poc-playground/bramble-federation/social/graph/generated"
)

const defaultPort = "4004"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	graph := generated.Config{
		Resolvers: &graph.Resolver{},
	}
	graph.Directives.Boundary = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		return next(ctx)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
