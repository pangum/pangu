package get

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/di"
)

type Commands struct {
	di.Get

	Commands []application.Command `group:"commands"`
}
