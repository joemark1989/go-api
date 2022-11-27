package bookInterface

import (
	"context"

	"github.com/go/crud/models"
)

type IBook interface {
	Books(ctx context.Context) ([]*models.Book, error)
	AddBook(ctx context.Context, input models.NewBook) (*models.Book, error)
}
