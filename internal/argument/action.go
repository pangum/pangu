package argument

import (
	"github.com/pangum/core/internal/constraint"
	"github.com/pangum/core/internal/runtime"
)

type Action[T constraint.Argument] func(ctx *runtime.Context, value T) error
