package runtime

import (
	"github.com/pangum/pangu/internal/argument"
	"github.com/pangum/pangu/internal/constraint"
)

// Argument 参数
type Argument[T constraint.Argument] struct {
	argument.Default[T]
}

// NewArgument 创建参数
func NewArgument[T constraint.Argument](name string, target *T) *argument.Builder[T] {
	return argument.New(name, target)
}
