package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*stringArg)(nil)

type stringArg struct {
	arg

	// 值
	value string
}

// NewString 创建一个字符串参数
func NewString(arg arg, value string) *stringArg {
	return &stringArg{
		arg:   arg,
		value: value,
	}
}

func (s *stringArg) Value() interface{} {
	return s.value
}

func (s *stringArg) ParseFlag() app.Flag {
	return &cli.StringFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       s.Value().(string),
		DefaultText: s.DefaultText(),
	}
}
