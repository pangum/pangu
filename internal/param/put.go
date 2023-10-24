package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	Constructors []runtime.Constructor
}

func NewPut(constructor runtime.Constructor, constructors ...runtime.Constructor) *Put {
	return &Put{
		Constructors: append([]runtime.Constructor{constructor}, constructors...),
	}
}
