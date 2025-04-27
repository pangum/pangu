package config

import (
	"github.com/harluo/boot/internal/internal/constant"
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
