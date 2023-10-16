package container

import (
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/storezhang/dig"
)

type Dependencies struct {
	container   *dig.Container
	params      *param.Dependencies
	constructor *verifier.Constructor
}

func NewDependencies(
	container *dig.Container,
	params *param.Dependencies,
	constructor *verifier.Constructor,
) *Dependencies {
	return &Dependencies{
		container:   container,
		params:      params,
		constructor: constructor,
	}
}

func (d *Dependencies) Put(constructors ...runtime.Constructor) (err error) {
	for _, constructor := range constructors {
		err = d.put(constructor)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependencies) Provide(constructors ...runtime.Constructor) {
	if err := d.Put(constructors...); nil != err {
		panic(err)
	}
}

func (d *Dependencies) Get(functions ...runtime.Getter) (err error) {
	for _, function := range functions {
		err = d.container.Invoke(function)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependencies) Invoke(functions ...runtime.Getter) {
	if err := d.Get(functions...); nil != err {
		panic(err)
	}
}

func (d *Dependencies) put(constructor runtime.Constructor) (err error) {
	if ve := d.constructor.Verify(constructor); nil != ve {
		err = ve
	} else if pe := d.container.Provide(constructor); nil != pe {
		err = pe
	}

	return
}
