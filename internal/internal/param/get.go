package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	Getter runtime.Getter
}

func NewGet(getter runtime.Getter) *Get {
	return &Get{
		Getter: getter,
	}
}
