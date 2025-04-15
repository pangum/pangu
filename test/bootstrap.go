package test

import (
	"github.com/pangum/core"
)

type Bootstrap struct {
	pangu.Lifecycle
}

func NewBootstrap() pangu.Bootstrap {
	return new(Bootstrap)
}

func (b *Bootstrap) Startup(_ *pangu.Application) (err error) {
	return
}
