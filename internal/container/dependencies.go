package container

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/storezhang/dig"
)

type Dependencies struct {
	container *dig.Container
	param     *param.Dependencies
}

func NewDependencies(container *dig.Container, param *param.Dependencies) *Dependencies {
	return &Dependencies{
		container: container,
		param:     param,
	}
}

func (d *Dependencies) Put(constructors ...any) (err error) {
	for _, constructor := range constructors {
		err = d.container.Provide(constructor)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependencies) Provide(constructors ...any) {
	if err := d.Put(constructors...); nil != err {
		panic(err)
	}
}

func (d *Dependencies) Get(functions ...any) (err error) {
	for _, function := range functions {
		err = d.container.Invoke(function)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependencies) Invoke(functions ...any) {
	if err := d.Get(functions...); nil != err {
		panic(err)
	}
}
