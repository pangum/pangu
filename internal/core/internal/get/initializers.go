package get

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/di"
)

type Initializers struct {
	di.Get

	Initializers []application.Initializer `group:"initializers"`
}
