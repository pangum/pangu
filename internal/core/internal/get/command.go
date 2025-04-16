package get

import (
	"github.com/heluon/boot/internal/container"
	"github.com/heluon/boot/internal/core/internal/command"
)

type Command struct {
	container.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
