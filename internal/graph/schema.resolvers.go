package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/subrotokumar/resonance/internal/model"
)

// Version is the resolver for the Version field.
func (r *mutationResolver) Version(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Version - Version"))
}

// Resonator is the resolver for the Resonator field.
func (r *queryResolver) Resonator(ctx context.Context, id *string, name *string) (*model.Resonator, error) {
	panic(fmt.Errorf("not implemented: Resonator - Resonator"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
