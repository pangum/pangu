package runtime

import (
	"github.com/pangum/pangu/internal/argument"
)

// Argument 参数
type Argument[T argument.Type] struct {
	argument.Default[T]
}

// NewArgument 创建参数
func NewArgument[T argument.Type](name string, target *T) *argument.Builder[T] {
	return argument.New(name, target)
}
