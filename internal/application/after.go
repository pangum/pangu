package application

import (
	"context"
)

type After interface {
	After(context.Context) error
}
