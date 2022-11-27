package bookInterface

import (
	"context"

	"github.com/go/crud/server/graph/model"
)

type IAuthor interface {
	Authors(ctx context.Context) ([]*model.Author, error)
}
