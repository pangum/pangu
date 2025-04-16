package argument

import (
	"github.com/heluon/boot/internal/constraint"
	"github.com/heluon/boot/internal/runtime"
)

type Action[T constraint.Argument] func(ctx *runtime.Context, value T) error
