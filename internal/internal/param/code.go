package param

import (
	"github.com/pangum/core/internal/internal/constant"
)

type Code struct {
	Success int
	Failed  int
	Panic   int
}

func newCode() *Code {
	return &Code{
		Success: constant.ApplicationCodeSuccess,
		Failed:  constant.ApplicationCodeFailed,
		Panic:   constant.ApplicationCodePanic,
	}
}
