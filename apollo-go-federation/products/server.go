//go:generate go run -mod=mod github.com/99designs/gqlgen
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/ihac/graphql-poc-playground/apollo-go-federation/common"
	"github.com/ihac/graphql-poc-playground/apollo-go-federation/products/graph"
)

const defaultPort = "4002"

var latencyInMills = flag.Int("l", 5, "additional latency")

func main() {
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.LatencyMiddleware(
		*latencyInMills,
		graph.GraphQLEndpointHandler(graph.EndpointOptions{EnableDebug: false, EnableRandomness: true}),
	))
	http.HandleFunc("/websocket_connections", graph.WebsocketConnectionsHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
