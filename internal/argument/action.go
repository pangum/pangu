package argument

import (
	"github.com/pangum/pangu/internal/internal/constraint"
	"github.com/pangum/pangu/internal/runtime"
)

type Action[T constraint.Argument] func(ctx *runtime.Context, value T) error
