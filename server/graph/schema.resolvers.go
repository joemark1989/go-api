package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/go/crud/models"
	"github.com/go/crud/server/graph/generated"
)

// AddBook is the resolver for the addBook field.
func (r *mutationResolver) AddBook(ctx context.Context, input models.NewBook) (*models.Book, error) {
	return r.Proxy.AddBook(ctx, input)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
	return r.Proxy.Books(ctx)
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*models.Author, error) {
	panic(fmt.Errorf("not implemented: Authors - authors"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
