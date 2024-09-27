package builder

import (
	"github.com/pangum/pangu/internal/container"
	"github.com/pangum/pangu/internal/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"go.uber.org/dig"
)

type Dependency struct {
	container *dig.Container
	params    *param.Dependency
}

func NewDependency(container *dig.Container, core *param.Application) *Dependency {
	return &Dependency{
		container: container,
		params:    param.NewDependency(core.Verify),
	}
}

func (d *Dependency) Invalidate() (dependency *Dependency) {
	d.params.Verify = false
	dependency = d

	return
}

func (d *Dependency) Verify() (dependency *Dependency) {
	d.params.Verify = true
	dependency = d

	return
}

func (d *Dependency) Put(constructor runtime.Constructor) *Put {
	return NewPut(d, constructor)
}

func (d *Dependency) Puts(required runtime.Constructor, others ...runtime.Constructor) (dependency *Dependency) {
	d.params.Puts = append(d.params.Puts, param.NewPut(required))
	for _, getter := range others {
		d.params.Puts = append(d.params.Puts, param.NewPut(getter))
	}
	dependency = d

	return
}

func (d *Dependency) Get(getter runtime.Getter) *Get {
	return NewGet(d, getter)
}

func (d *Dependency) Gets(required runtime.Getter, others ...runtime.Getter) (dependency *Dependency) {
	d.params.Gets = append(d.params.Gets, param.NewGet(required))
	for _, getter := range others {
		d.params.Gets = append(d.params.Gets, param.NewGet(getter))
	}
	dependency = d

	return
}

func (d *Dependency) Build() *container.Dependency {
	return container.NewDependency(d.container, d.params)
}
