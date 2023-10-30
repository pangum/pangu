package app

import (
	"context"
)

// Stopper 可停止
type Stopper interface {
	// Stop 停止
	Stop(ctx context.Context) (err error)
}
