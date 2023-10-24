package param

import (
	"github.com/pangum/pangu/internal/constant"
)

type Code struct {
	Success int
	Failed  int
	Panic   int
	Set     bool
}

func newCode() *Code {
	return &Code{
		Success: constant.ApplicationCodeSuccess,
		Failed:  constant.ApplicationCodeFailed,
		Panic:   constant.ApplicationCodePanic,
	}
}
