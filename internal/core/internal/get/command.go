package get

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/core/internal/command"
)

type Command struct {
	container.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
