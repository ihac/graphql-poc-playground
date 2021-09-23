package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ihac/graphql-poc-playground/apollo-go-federation/accounts/graph/generated"
	"github.com/ihac/graphql-poc-playground/apollo-go-federation/accounts/graph/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	fmt.Println("[accounts] Me()")
	fmt.Println("\t>> select * from users limit 1")
	return &model.User{
		ID:       "1234",
		Username: "Me",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
