package boot

import (
	"github.com/harluo/boot/internal/builder"
	"github.com/harluo/boot/internal/core"
)

// Application 应用程序
type Application = core.Application

func New() *builder.Application {
	return builder.NewApplication()
}
