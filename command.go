package boot

import (
	"github.com/harluo/boot/internal/builder"
)

// NewCommand 创建命令
func NewCommand(name string) *builder.Command {
	return builder.NewCommand(name)
}
