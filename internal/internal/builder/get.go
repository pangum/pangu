package builder

import (
	"github.com/pangum/pangu/internal/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	dependency *Dependency
	params     *param.Get
}

func NewGet(dependency *Dependency, getter runtime.Getter) *Get {
	return &Get{
		dependency: dependency,
		params:     param.NewGet(getter),
	}
}

func (g *Get) Build() (dependency *Dependency) {
	g.dependency.params.Gets = append(g.dependency.params.Gets, g.params)
	dependency = g.dependency

	return
}
