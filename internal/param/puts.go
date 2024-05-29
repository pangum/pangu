package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Puts struct {
	Constructors []runtime.Constructor
}

func NewPuts(constructor runtime.Constructor, constructors ...runtime.Constructor) *Puts {
	return &Puts{
		Constructors: append([]runtime.Constructor{constructor}, constructors...),
	}
}
