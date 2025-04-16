package config

import (
	"github.com/goexl/log"
	"github.com/heluon/boot/internal/internal/param"
	"github.com/heluon/boot/internal/runtime"
)

type Setup struct {
	getter *Getter
}

func NewSetup(params *param.Config, logger *log.Logger) *Setup {
	return &Setup{
		getter: newGetter(params, logger),
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
