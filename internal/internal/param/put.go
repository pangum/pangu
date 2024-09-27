package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Put struct {
	Constructor runtime.Constructor
	Name        string
	Group       string
}

func NewPut(constructor runtime.Constructor) *Put {
	return &Put{
		Constructor: constructor,
	}
}
