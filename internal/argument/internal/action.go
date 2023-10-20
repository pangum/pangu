package internal

import (
	"github.com/pangum/pangu/internal/argument"
	"github.com/pangum/pangu/internal/runtime"
)

type Action[T argument.Type] func(ctx *runtime.Context, value T) error
