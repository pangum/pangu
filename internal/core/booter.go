package core

import (
	"context"
)

// Booter 启动器，全局只能有一个启动器，且只能返回 Booter 才能被正确的启动
type Booter interface {
	// Boot 启动
	Boot(ctx context.Context) error
}
