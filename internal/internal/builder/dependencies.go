package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/internal/verifier"
	"github.com/pangum/pangu/internal/param"
	"github.com/storezhang/dig"
)

type Dependencies struct {
	container   *dig.Container
	params      *param.Dependencies
	constructor *verifier.Constructor
}

func NewDependencies(container *dig.Container, core *param.Application) *Dependencies {
	return &Dependencies{
		container:   container,
		constructor: verifier.NewConstructor(core),
	}
}

func (d *Dependencies) Build() *container.Dependencies {
	return container.NewDependencies(d.container, d.params, d.constructor)
}
