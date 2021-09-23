package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ihac/graphql-poc-playground/apollo-go-federation/accounts/graph/generated"
	"github.com/ihac/graphql-poc-playground/apollo-go-federation/accounts/graph/model"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	fmt.Printf("[accounts] FindUserByID(%s)\n", id)
	fmt.Printf("\t>> select * from users where id = %s\n", id)
	name := "User " + id
	if id == "1234" {
		name = "Me"
	}

	return &model.User{
		ID:       id,
		Username: name,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
