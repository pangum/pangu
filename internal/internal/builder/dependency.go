package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
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
	}
}

func (d *Dependency) Build() *container.Dependency {
	return container.NewDependency(d.container, d.params, d.constructor)
}
