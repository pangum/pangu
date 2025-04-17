package argument

import (
	"github.com/harluo/boot/internal/constraint"
	"github.com/harluo/boot/internal/runtime"
)

type Action[T constraint.Argument] func(ctx *runtime.Context, value T) error
