package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/ihac/graphql-poc-playground/bramble-federation/masthead/graph/generated"
	"github.com/ihac/graphql-poc-playground/bramble-federation/masthead/graph/model"
)

func (r *queryResolver) Masthead(ctx context.Context, count *int) ([]model.MastheadItem, error) {
	fmt.Printf("query: masthead %d\n", *count)
	var retval []model.MastheadItem
	for i := 0; i < *count; i++ {
		typeIdx := rand.Intn(3)
		id := strconv.Itoa(rand.Intn(1000))
		switch typeIdx {
		case 0:
			retval = append(retval, &model.ContentItem{ID: id})
		case 1:
			retval = append(retval, &model.AdvtItem{ID: id})
		case 2:
			retval = append(retval, &model.ScoreCard{ID: id})
		}
	}
	return retval, nil
}

func (r *queryResolver) ContentItem(ctx context.Context, id string) (*model.ContentItem, error) {
	fmt.Printf("query: content %s\n", id)
	return &model.ContentItem{ID: id}, nil
}

func (r *queryResolver) AdvtItem(ctx context.Context, id string) (*model.AdvtItem, error) {
	fmt.Printf("query: ad %s\n", id)
	return &model.AdvtItem{ID: id}, nil
}

func (r *queryResolver) ScoreCard(ctx context.Context, id string) (*model.ScoreCard, error) {
	fmt.Printf("query: score card %s\n", id)
	return &model.ScoreCard{ID: id}, nil
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

func (r *queryResolver) RandomContent(ctx context.Context) (*model.ContentItem, error) {
	fmt.Println("query: random")
	return &model.ContentItem{ID: "random"}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
