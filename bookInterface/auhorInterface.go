package bookinterface

import (
	"context"

	"github.com/go/crud/bookServer/gqlgen/graph/model"
)

type IAuthor interface {
	Authors(ctx context.Context) ([]*model.Author, error)
}
