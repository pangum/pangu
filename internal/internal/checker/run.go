package checker

import (
	"context"
)

type Run interface {
	Run(context.Context) error
}
