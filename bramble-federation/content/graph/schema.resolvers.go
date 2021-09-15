package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/ihac/graphql-poc-playground/bramble-federation/content/graph/generated"
	"github.com/ihac/graphql-poc-playground/bramble-federation/content/graph/model"
)

func (r *queryResolver) ContentItem(ctx context.Context, id string) (*model.ContentItem, error) {
	fmt.Printf("query: %s\n", id)
	idx := rand.Intn(len(titles))
	return &model.ContentItem{
		ID:          id,
		Title:       titles[idx],
		Description: fmt.Sprintf("This is the description of %s", titles[idx]),
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

var (
	titles = []string{
		"Traitor Of Sunshine",
		"Friend In The News",
		"Recruit Of The Ocean",
		"Mercenary Of Time",
		"Intruders From The UFO",
		"Hunters Of The Orbit",
		"Directors Of The Void",
		"Medics Of Outer Space",
		"Invaders And Heroes",
		"Captains And Captains",
		"Enemies And Heroes",
		"Volunteers And Invaders",
		"Victory Of Our Ship",
		"Result",
		"Death Of The Universe",
		"Restoration Of Everywhere",
		"Married To My Planet",
		"Lonely In The Immortals",
		"Fortune Of The Depths",
		"Lonely In Eternity",
		"Anxious For Robotic Control",
		"Father Of Technolic Advancements",
		"Fortune Of My Journey",
		"Equality Of Moon Rocks}",
	}
)
