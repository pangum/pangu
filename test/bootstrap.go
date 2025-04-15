package test

import (
	"github.com/pangum/core"
)

type Bootstrap struct {
	core.Lifecycle
}

func NewBootstrap() core.Bootstrap {
	return new(Bootstrap)
}

func (b *Bootstrap) Startup(_ *core.Application) (err error) {
	return
}
