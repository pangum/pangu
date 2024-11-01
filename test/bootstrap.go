package test

import (
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Lifecycle
}

func NewBootstrap() *Bootstrap {
	return new(Bootstrap)
}

func (b *Bootstrap) Startup(_ *pangu.Application) (err error) {
	return
}
