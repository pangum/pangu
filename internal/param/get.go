package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	Injection *Injection
	Getter    runtime.Getter
}

func NewGet(getter runtime.Getter) *Get {
	return &Get{
		Injection: NewInjection(),
		Getter:    getter,
	}
}
