package boot

import (
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/argument"
	"github.com/harluo/boot/internal/constraint"
)

// Argument 参数
type Argument = application.Argument

// NewArgument 创建参数
func NewArgument[T constraint.Argument](name string, target *T) *argument.Builder[T] {
	return argument.New(name, target)
}
