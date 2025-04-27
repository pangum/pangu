package application

import (
	"context"
)

type After interface {
	After(ctx context.Context) error
}
