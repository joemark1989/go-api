package proxy

import (
	"github.com/go/crud/bookInterface"
	"gorm.io/gorm"
)

type Proxy struct {
	DB *gorm.DB
	bookInterface.IProxy
}
