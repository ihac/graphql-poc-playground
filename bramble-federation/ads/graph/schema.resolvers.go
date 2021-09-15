package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"

	"github.com/ihac/graphql-poc-playground/bramble-federation/ads/graph/generated"
	"github.com/ihac/graphql-poc-playground/bramble-federation/ads/graph/model"
)

func (r *queryResolver) AdvtItem(ctx context.Context, id string) (*model.AdvtItem, error) {
	fmt.Printf("query: %s\n", id)
	return &model.AdvtItem{
		ID:         id,
		TrackerURL: "https://example.com/tracker/url",
	}, nil
}

func (r *queryResolver) Service(ctx context.Context) (*model.Service, error) {
	fmt.Println("query: Service")
	content, err := os.ReadFile("./graph/schema.graphqls")
	if err != nil {
		return nil, fmt.Errorf("unable to read the schema: %v", err)
	}

	return &model.Service{
		Name:    "masthead",
		Version: "1.0.1",
		Schema:  string(content),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
