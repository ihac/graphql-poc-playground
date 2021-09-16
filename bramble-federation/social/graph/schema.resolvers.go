package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/ihac/graphql-poc-playground/bramble-federation/social/graph/generated"
	"github.com/ihac/graphql-poc-playground/bramble-federation/social/graph/model"
)

func (r *queryResolver) ScoreCard(ctx context.Context, id string) (*model.ScoreCard, error) {
	fmt.Printf("query: score card %s\n", id)
	score := rand.Intn(50) + 50
	return &model.ScoreCard{
		ID:         id,
		TeamAScore: score + rand.Intn(20),
		TeamBScore: score + rand.Intn(20),
	}, nil
}

func (r *queryResolver) ScoreCardAds(ctx context.Context, id string) (*model.TinyAdsBanner, error) {
	fmt.Printf("query: ads %s\n", id)
	return &model.TinyAdsBanner{ID: id}, nil
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
