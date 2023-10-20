package runtime

import (
	"github.com/pangum/pangu/internal/command"
)

// Command 命令
type Command = command.Default

// NewCommand 创建命令
func NewCommand(name string) *command.Builder {
	return command.New(name)
}
