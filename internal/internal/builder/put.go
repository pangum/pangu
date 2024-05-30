package builder

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	dependency *Dependency
	params     *param.Put
}

func NewPut(dependency *Dependency, constructor runtime.Constructor) (put *Put) {
	return &Put{
		dependency: dependency,
		params:     param.NewPut(constructor),
	}
}

func (p *Put) Name(name string) (put *Put) {
	p.params.Name = name
	put = p

	return
}

func (p *Put) Group(group string) (put *Put) {
	p.params.Group = group
	put = p

	return
}

func (p *Put) Build() (dependency *Dependency) {
	p.dependency.params.Puts = append(p.dependency.params.Puts, p.params)
	dependency = p.dependency

	return
}
