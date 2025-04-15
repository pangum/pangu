package builder

import (
	"github.com/pangum/core/internal/internal/param"
	"github.com/pangum/core/internal/runtime"
)

type Put struct {
	dependency *Dependency
	params     *param.Put
}

func NewPut(dependency *Dependency, constructor runtime.Constructor) *Put {
	return &Put{
		dependency: dependency,
		params:     param.NewPut(constructor),
	}
}

func (p *Put) Name(name string, names ...string) (put *Put) {
	p.params.Names = append(p.params.Names, append([]string{name}, names...)...)
	put = p

	return
}

func (p *Put) Group(group string, groups ...string) (put *Put) {
	p.params.Groups = append(p.params.Groups, append([]string{group}, groups...)...)
	put = p

	return
}

func (p *Put) Build() (dependency *Dependency) {
	p.dependency.params.Puts = append(p.dependency.params.Puts, p.params)
	dependency = p.dependency

	return
}
