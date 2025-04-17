package get

import (
	"github.com/harluo/boot/internal/core/internal/command"
	"github.com/harluo/di"
)

type Command struct {
	di.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
