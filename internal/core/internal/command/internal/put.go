package internal

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/di"
)

type Put struct {
	di.Put

	Info    application.Command `group:"commands"`
	Version application.Command `group:"commands"`
}
