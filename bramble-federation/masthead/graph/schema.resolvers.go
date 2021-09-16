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

func (r *queryResolver) MastheadV2(ctx context.Context, count *int) (*model.Masthead, error) {
	fmt.Printf("query: masthead v2 %d\n", *count)

	masthead := model.Masthead{}
	rem := *count
	for i := 0; i < 4; i++ {
		cnt := rand.Intn(*count / 2)
		if cnt > rem {
			cnt = rem
		}

		if i == 3 {
			cnt = rem
		}

		switch i {
		case 0:
			for j := 0; j < cnt; j++ {
				masthead.ContentItems = append(masthead.ContentItems, &model.ContentItem{ID: strconv.Itoa(rand.Intn(1000))})
			}
		case 1:
			for j := 0; j < cnt; j++ {
				masthead.ContentItemsWithScoreCard = append(masthead.ContentItemsWithScoreCard, &model.ContentItemWithScoreCard{ID: strconv.Itoa(rand.Intn(1000))})
			}
		case 2:
			for j := 0; j < cnt; j++ {
				masthead.ScoreCards = append(masthead.ScoreCards, &model.ScoreCard{ID: strconv.Itoa(rand.Intn(1000))})
			}
		case 3:
			for j := 0; j < cnt; j++ {
				masthead.AdvtItems = append(masthead.AdvtItems, &model.AdvtItem{ID: strconv.Itoa(rand.Intn(1000))})
			}
		}
		rem -= cnt
	}
	return &masthead, nil
}

func (r *queryResolver) ContentItem(ctx context.Context, id string) (*model.ContentItem, error) {
	fmt.Printf("query: content %s\n", id)
	return &model.ContentItem{ID: id}, nil
}

func (r *queryResolver) ContentItemWithScoreCard(ctx context.Context, id string) (*model.ContentItemWithScoreCard, error) {
	fmt.Printf("query: content with score %s\n", id)
	return &model.ContentItemWithScoreCard{ID: id}, nil
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) RandomContent(ctx context.Context) (*model.ContentItem, error) {
	fmt.Println("query: random")
	return &model.ContentItem{ID: "random"}, nil
}
