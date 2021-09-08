package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hotstar/graphql-poc-playground/graph/generated"
	"github.com/hotstar/graphql-poc-playground/graph/model"
)

func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	return r.InsertCompany(&input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.InsertUser(&input)
}

func (r *queryResolver) Companies(ctx context.Context) ([]*model.Company, error) {
	return r.ListAllCompanies()
}

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	return r.GetCompanyByID(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.ListAllUsers()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.GetUserByID(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
