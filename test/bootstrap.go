package test

import (
	"github.com/pangum/pangu"
)

type Bootstrap struct {
	pangu.Bootstrap
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) Startup() (err error) {
	return
}
