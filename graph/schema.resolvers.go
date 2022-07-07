package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/FVitor7/golang-graphql-server/graph/generated"
	"github.com/FVitor7/golang-graphql-server/graph/model"
)

// CreateBrand is the resolver for the createBrand field.
func (r *mutationResolver) CreateBrand(ctx context.Context, input model.NewBrand) (*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateCar is the resolver for the createCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input model.NewCar) (*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

// Brands is the resolver for the brands field.
func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
