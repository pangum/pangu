package runtime

import (
	argument2 "github.com/pangum/pangu/internal/internal/argument"
)

// Argument 参数
type Argument[T argument2.Type] struct {
	argument2.Default[T]
}

// NewArgument 创建参数
func NewArgument[T argument2.Type](name string, target *T) *argument2.Builder[T] {
	return argument2.New(name, target)
}
