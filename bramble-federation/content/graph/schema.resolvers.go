package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/ihac/graphql-poc-playground/bramble-federation/content/graph/generated"
	"github.com/ihac/graphql-poc-playground/bramble-federation/content/graph/model"
)

func (r *queryResolver) ContentItem(ctx context.Context, id string) (*model.ContentItem, error) {
	fmt.Printf("query: content %s\n", id)
	idx := rand.Intn(len(titles))
	return &model.ContentItem{
		ID:          id,
		Title:       titles[idx],
		Description: fmt.Sprintf("This is the description of %s", titles[idx]),
	}, nil
}

func (r *queryResolver) ContentItemWithScoreCard(ctx context.Context, id string) (*model.ContentItemWithScoreCard, error) {
	fmt.Printf("query: content with score %s\n", id)
	idx := rand.Intn(len(titles))
	return &model.ContentItemWithScoreCard{
		ID:    id,
		Title: titles[idx],
		Score: &model.ScoreCard{
			ID: strconv.Itoa(rand.Intn(1000)),
		},
	}, nil
}

func (r *queryResolver) ContentItemWithScoreCardScore(ctx context.Context, id string) (*model.ScoreCard, error) {
	fmt.Printf("query: score card %s\n", id)
	return &model.ScoreCard{ID: id}, nil
}

func (r *queryResolver) RandomContent(ctx context.Context) (*model.ContentItem, error) {
	fmt.Println("query: random content")
	idx := rand.Intn(len(titles))
	return &model.ContentItem{
		ID:          "random",
		Title:       titles[idx],
		Description: fmt.Sprintf("This is the description of %s", titles[idx]),
	}, nil
}

func (r *queryResolver) RandomContentWithScoreCard(ctx context.Context) (*model.ContentItemWithScoreCard, error) {
	fmt.Println("query: random content with score")
	idx := rand.Intn(len(titles))
	return &model.ContentItemWithScoreCard{
		ID:    "random",
		Title: titles[idx],
		Score: &model.ScoreCard{
			ID: strconv.Itoa(rand.Intn(1000)),
		},
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
