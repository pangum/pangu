package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*stringsArg)(nil)

type stringsArg struct {
	arg

	// 值
	value []string
}

// NewStringSlice 创建一个字符串数组参数
func NewStringSlice(arg arg, values ...string) *stringsArg {
	return &stringsArg{
		arg:   arg,
		value: values,
	}
}

func (s *stringsArg) Value() interface{} {
	return s.value
}

func (s *stringsArg) ParseFlag() app.Flag {
	return &cli.StringSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewStringSlice(s.Value().([]string)...),
		DefaultText: s.DefaultText(),
	}
}
