package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/storezhang/dig"
)

type Dependency struct {
	container   *dig.Container
	constructor *verifier.Constructor
	params      *param.Dependency
}

func NewDependency(container *dig.Container, core *param.Application) *Dependency {
	return &Dependency{
		constructor: verifier.NewConstructor(core),
		container:   container,
		params:      param.NewDependency(),
	}
}

func (d *Dependency) Put(constructor runtime.Constructor, constructors ...runtime.Constructor) *Put {
	return NewPut(d, constructor, constructors...)
}

func (d *Dependency) Get(getter runtime.Getter, getters ...runtime.Getter) *Get {
	return NewGet(d, getter, getters...)
}

func (d *Dependency) Build() *container.Dependency {
	return container.NewDependency(d.container, d.params, d.constructor)
}
