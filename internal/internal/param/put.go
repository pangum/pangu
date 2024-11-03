package param

import (
	"github.com/pangum/pangu/internal/internal/constant"
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	Constructor runtime.Constructor
	Names       []string
	Groups      []string
}

func NewPut(constructor runtime.Constructor) *Put {
	return &Put{
		Constructor: constructor,
		Names:       []string{constant.DependencyNone},
		Groups:      []string{constant.DependencyNone},
	}
}
