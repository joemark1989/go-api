package graph

import (
	bookinterface "github.com/go/crud/bookInterface"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Proxy bookinterface.IProxy
}
