package get

import (
	"github.com/harluo/boot/internal/container"
	"github.com/harluo/boot/internal/core/internal/command"
)

type Command struct {
	container.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
