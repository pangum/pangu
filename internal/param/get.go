package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Get struct {
	Getters []runtime.Getter
}

func NewGet(getter runtime.Getter, getters ...runtime.Getter) *Get {
	return &Get{
		Getters: append([]runtime.Getter{getter}, getters...),
	}
}
