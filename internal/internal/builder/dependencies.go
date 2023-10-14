package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/param"
	"github.com/storezhang/dig"
)

type Dependencies struct {
	container *dig.Container
	param     *param.Dependencies
}

func NewDependencies(container *dig.Container) *Dependencies {
	return &Dependencies{
		container: container,
	}
}

func (d *Dependencies) Build() *container.Dependencies {
	return container.NewDependencies(d.container, d.param)
}
