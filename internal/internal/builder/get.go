package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	core *Dependency
	gets []*param.Get
}

func NewGet(core *Dependency, getter runtime.Getter, getters ...runtime.Getter) *Get {
	return &Get{
		core: core,
		gets: []*param.Get{
			param.NewGet(getter, getters...),
		},
	}
}

func (g *Get) Build() (dependency *Dependency) {
	g.core.params.Gets = append(g.core.params.Gets, g.gets...)
	dependency = g.core

	return
}
