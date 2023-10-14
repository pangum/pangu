package container

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/storezhang/dig"
)

type Dependency struct {
	container *dig.Container
	param     *param.Dependency
}

func NewDependency(container *dig.Container, param *param.Dependency) *Dependency {
	return &Dependency{
		container: container,
		param:     param,
	}
}

func (d *Dependency) Put(constructor any) error {
	return d.container.Provide(constructor)
}

func (d *Dependency) Provide(constructor any) {
	if err := d.Put(constructor); nil != err {
		panic(err)
	}
}

func (d *Dependency) Get(function any) error {
	return d.container.Invoke(function)
}

func (d *Dependency) Invoke(function any) {
	if err := d.Get(function); nil != err {
		panic(err)
	}
}
