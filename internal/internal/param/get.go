package param

import (
	"github.com/pangum/core/internal/runtime"
)

type Get struct {
	Getter runtime.Getter
}

func NewGet(getter runtime.Getter) *Get {
	return &Get{
		Getter: getter,
	}
}
