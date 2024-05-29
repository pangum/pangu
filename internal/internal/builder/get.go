package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	*Injection

	dependency *Dependency
	params     *param.Get
}

func NewGet(dependency *Dependency, getter runtime.Getter) (get *Get) {
	get = new(Get)
	get.dependency = dependency
	get.params = param.NewGet(getter)
	get.Injection = NewInjection(get.params.Injection)

	return
}

func (g *Get) Build() (dependency *Dependency) {
	g.dependency.params.Gets = append(g.dependency.params.Gets, g.params)
	dependency = g.dependency

	return
}
