package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ihac/graphql-poc-playground/apollo-go-federation/reviews/graph/generated"
	"github.com/ihac/graphql-poc-playground/apollo-go-federation/reviews/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	fmt.Printf("[reviews] FindProductByUpc(%s)\n", upc)
	fmt.Printf("\t>> select product from reviews where upc = %s\n", upc)
	return &model.Product{
		Upc: upc,
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	fmt.Printf("[reviews] FindUserByID(%s)\n", id)
	fmt.Printf("\t>> select user from reviews where user_id = %s\n", id)
	return &model.User{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
