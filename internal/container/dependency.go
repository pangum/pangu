package container

import (
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/storezhang/dig"
)

type Dependency struct {
	container   *dig.Container
	params      *param.Dependency
	constructor *verifier.Constructor
}

func NewDependency(container *dig.Container, params *param.Dependency, constructor *verifier.Constructor) *Dependency {
	return &Dependency{
		container:   container,
		constructor: constructor,
		params:      params,
	}
}

func (d *Dependency) Inject() (err error) {
	if de := d.puts(); nil != de {
		err = de
	} else if ge := d.gets(); nil != ge {
		err = ge
	}

	return
}

func (d *Dependency) Apply() {
	if err := d.Inject(); nil != err {
		panic(err)
	}
}

func (d *Dependency) gets() (err error) {
	for _, get := range d.params.Gets {
		err = d.invoke(get)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) invoke(get *param.Get) (err error) {
	for _, getter := range get.Getters {
		err = d.container.Invoke(getter)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) puts() (err error) {
	for _, put := range d.params.Puts {
		err = d.put(put)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) put(put *param.Put) (err error) {
	for _, constructor := range put.Constructors {
		err = d.provide(constructor)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) provide(constructor runtime.Constructor) (err error) {
	if ve := d.constructor.Verify(constructor); nil != ve {
		err = ve
	} else if pe := d.container.Provide(constructor); nil != pe {
		err = pe
	}

	return
}
