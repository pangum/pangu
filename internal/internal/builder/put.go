package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	core *Dependency
	puts []*param.Put
}

func NewPut(core *Dependency, constructor runtime.Constructor, constructors ...runtime.Constructor) (put *Put) {
	return &Put{
		core: core,
		puts: []*param.Put{
			param.NewPut(constructor, constructors...),
		},
	}
}

func (p *Put) Build() (dependency *Dependency) {
	p.core.params.Puts = append(p.core.params.Puts, p.puts...)
	dependency = p.core

	return
}
