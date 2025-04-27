package boot

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/builder"
)

// Command 命令
type Command = application.Command

// NewCommand 创建命令
func NewCommand(name string) *builder.Command {
	return builder.NewCommand(name)
}
