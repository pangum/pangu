package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/param"
	"github.com/storezhang/dig"
)

type Dependency struct {
	container *dig.Container
	param     *param.Dependency
}

func NewDependency(container *dig.Container) *Dependency {
	return &Dependency{
		container: container,
	}
}

func (d *Dependency) Build() *container.Dependency {
	return container.NewDependency(d.container, d.param)
}
