package get

import (
	"github.com/pangum/pangu/internal/command"
	"github.com/pangum/pangu/internal/runtime"
)

type Command struct {
	runtime.Get

	Serve   *command.Serve
	Info    *command.Info
	Version *command.Version
}
