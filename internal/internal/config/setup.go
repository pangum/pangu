package config

import (
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Setup struct {
	getter *Getter
}

func NewSetup(params *param.Config) *Setup {
	return &Setup{
		getter: newGetter(params),
	}
}

func (s *Setup) Bind(shell *runtime.Shell, shadow *runtime.Shadow) {
	s.getter.bind(shell, shadow)
}

func (s *Setup) Copy(params *param.Config) (config *Getter) {
	s.getter.params = params
	config = s.getter

	return
}
