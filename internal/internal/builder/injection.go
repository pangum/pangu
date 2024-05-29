package builder

import (
	"github.com/pangum/pangu/internal/param"
)

type Injection struct {
	params *param.Injection
}

func NewInjection(params *param.Injection) *Injection {
	return &Injection{
		params: params,
	}
}

func (i *Injection) Name(name string) (injection *Injection) {
	i.params.Name = name
	injection = i

	return
}

func (i *Injection) Group(group string) (injection *Injection) {
	i.params.Group = group
	injection = i

	return
}
