package boot

import (
	"github.com/harluo/boot/internal/builder"
	"github.com/harluo/boot/internal/core"
)

// Command 命令
type Command = core.Command

// NewCommand 创建命令
func NewCommand(name string) *builder.Command {
	return builder.NewCommand(name)
}
