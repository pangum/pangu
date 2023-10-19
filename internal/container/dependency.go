package container

import (
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/storezhang/dig"
)

type Dependency struct {
	container   *dig.Container
	param       *param.Dependency
	constructor *verifier.Constructor
}

func NewDependency(container *dig.Container, params *param.Dependency, constructor *verifier.Constructor) *Dependency {
	return &Dependency{
		container:   container,
		constructor: constructor,
		param:       params,
	}
}

func (d *Dependency) Put(constructor runtime.Constructor) (err error) {
	if ve := d.constructor.Verify(constructor); nil != ve {
		err = ve
	} else if pe := d.container.Provide(constructor); nil != pe {
		err = pe
	}

	return
}

func (d *Dependency) Provide(constructor runtime.Constructor) {
	if err := d.Put(constructor); nil != err {
		panic(err)
	}
}

func (d *Dependency) Get(getter runtime.Getter) error {
	return d.container.Invoke(getter)
}

func (d *Dependency) Invoke(getter runtime.Getter) {
	if err := d.Get(getter); nil != err {
		panic(err)
	}
}
