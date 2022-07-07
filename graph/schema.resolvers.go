package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/FVitor7/golang-graphql-server/graph/generated"
	"github.com/FVitor7/golang-graphql-server/graph/model"
)

// Cars is the resolver for the cars field.
func (r *brandResolver) Cars(ctx context.Context, obj *model.Brand) ([]*model.Car, error) {
	var cars []*model.Car

	for _, v := range r.Resolver.Cars {
		if v.Brand.ID == obj.ID {
			cars = append(cars, v)
		}
	}

	return cars, nil
}

// CreateBrand is the resolver for the createBrand field.
func (r *mutationResolver) CreateBrand(ctx context.Context, input model.NewBrand) (*model.Brand, error) {
	brand := model.Brand{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}
	r.Brands = append(r.Brands, &brand)

	return &brand, nil
}

// CreateCar is the resolver for the createCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input model.NewCar) (*model.Car, error) {
	var brand *model.Brand

	for _, v := range r.Brands {
		if v.ID == input.BrandID {
			brand = v
		}
	}

	car := model.Car{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Name:  input.Name,
		Brand: brand,
	}

	r.Cars = append(r.Cars, &car)
	return &car, nil
}

// Brands is the resolver for the brands field.
func (r *queryResolver) Brands(ctx context.Context) ([]*model.Brand, error) {
	return r.Resolver.Brands, nil
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	return r.Resolver.Cars, nil
}

// Brand returns generated.BrandResolver implementation.
func (r *Resolver) Brand() generated.BrandResolver { return &brandResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type brandResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
