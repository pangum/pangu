package application

import (
	"context"
)

type Stopper interface {
	Stop(ctx context.Context) error
}
