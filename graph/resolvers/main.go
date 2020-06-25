package resolvers

import (
	"gqlgen-todos/graph/generated"
	"gqlgen-todos/graph/model"
)

type Resolver struct {
	authors []*model.Author
	books   []*model.Book
	author  *model.Author
	book    *model.Book
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
