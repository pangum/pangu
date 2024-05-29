package param

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Gets struct {
	Getters []runtime.Getter
}

func NewGets(getter runtime.Getter, getters ...runtime.Getter) *Gets {
	return &Gets{
		Getters: append([]runtime.Getter{getter}, getters...),
	}
}
