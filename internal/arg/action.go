package arg

import (
	"github.com/pangum/pangu/internal/app"
)

type action[T argumentType] func(ctx *app.Context, value T) error
