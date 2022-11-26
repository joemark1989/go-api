package bookinterface

import (
	"context"

	"github.com/go/crud/bookServer/gqlgen/graph/model"
)

type IBook interface {
	Books(ctx context.Context) ([]*model.Book, error)
	AddBook(ctx context.Context, input model.NewBook) (*model.Book, error)
}
