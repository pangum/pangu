package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	dependency *Dependency
	gets       []*param.Get
}

func NewGet(dependency *Dependency, getter runtime.Getter, getters ...runtime.Getter) *Get {
	return &Get{
		dependency: dependency,
		gets: []*param.Get{
			param.NewGet(getter, getters...),
		},
	}
}

func (g *Get) Build() (dependency *Dependency) {
	g.dependency.params.Gets = append(g.dependency.params.Gets, g.gets...)
	dependency = g.dependency

	return
}
