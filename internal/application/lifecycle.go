package application

import (
	"context"
)

type Lifecycle interface {
	// Before 程序启动前
	Before(ctx context.Context) error

	// After 程序启动后
	After(ctx context.Context) error
}
