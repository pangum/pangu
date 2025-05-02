package application

import (
	"context"
)

// Command 描述一个可以被执行的命令
type Command interface {
	Parameter

	// Run 执行命令
	Run(context.Context) error
}
