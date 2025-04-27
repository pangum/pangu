package get

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/di"
)

type Arguments struct {
	di.Get

	Arguments []application.Argument `group:"arguments"`
}
