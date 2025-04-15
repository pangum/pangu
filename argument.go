package core

import (
	"github.com/pangum/core/internal/argument"
	"github.com/pangum/core/internal/constraint"
)

// Argument 参数
type Argument[T constraint.Argument] struct {
	argument.Default[T]
}

// NewArgument 创建参数
func NewArgument[T constraint.Argument](name string, target *T) *argument.Builder[T] {
	return argument.New(name, target)
}
