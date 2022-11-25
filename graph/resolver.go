package graph

import "github.com/go/crud/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookStore map[string]model.Book
}
