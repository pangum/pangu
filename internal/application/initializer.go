package application

import (
	"context"
)

type Initializer interface {
	Initialize(ctx context.Context) error
}
