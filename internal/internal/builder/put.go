package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	dependency *Dependency
	puts       []*param.Put
}

func NewPut(dependency *Dependency, constructor runtime.Constructor, constructors ...runtime.Constructor) (put *Put) {
	return &Put{
		dependency: dependency,
		puts: []*param.Put{
			param.NewPut(constructor, constructors...),
		},
	}
}

func (p *Put) Build() (dependency *Dependency) {
	p.dependency.params.Puts = append(p.dependency.params.Puts, p.puts...)
	dependency = p.dependency

	return
}
