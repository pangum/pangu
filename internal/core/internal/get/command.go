package get

import (
	"github.com/pangum/core/internal/container"
	"github.com/pangum/core/internal/core/internal/command"
)

type Command struct {
	container.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
