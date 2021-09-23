package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ihac/graphql-poc-playground/apollo-go-federation/products/graph/generated"
	"github.com/ihac/graphql-poc-playground/apollo-go-federation/products/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	fmt.Printf("[products] FindProductByUpc(%s)\n", upc)
	fmt.Printf("\t>> select * from products where upc = %s\n", upc)
	for _, h := range hats {
		if h.Upc == upc {
			return h, nil
		}
	}
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
