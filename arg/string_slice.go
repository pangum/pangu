package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*stringsArg)(nil)

type stringsArg struct {
	*base

	// 值
	value []string
}

// NewStringSlice 创建一个字符串数组参数
func NewStringSlice(base *base, values ...string) *stringsArg {
	return &stringsArg{
		base:  base,
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
