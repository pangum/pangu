package argument

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Action[T Type] func(ctx *runtime.Context, value T) error
