package application

import (
	"context"
)

type Before interface {
	Before(ctx context.Context) error
}
