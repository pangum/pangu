package get

import (
	"github.com/pangum/pangu/internal/container"
	command2 "github.com/pangum/pangu/internal/internal/command"
)

type Command struct {
	container.Get

	Serve   *command2.Serve
	Info    *command2.Info
	Version *command2.Version
}
