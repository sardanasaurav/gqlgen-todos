package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
)

type Resolver struct{}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Book(ctx context.Context, id *string) (*model.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Author(ctx context.Context, id *string) (*model.Author, error) {
	panic("not implemented")
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic("not implemented")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
