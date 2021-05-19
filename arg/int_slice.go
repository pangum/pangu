package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*intSliceArg)(nil)

type intSliceArg struct {
	*base

	// 值
	value []int
}

// NewIntSlice 创建一个整形数组参数
func NewIntSlice(base *base, values ...int) *intSliceArg {
	return &intSliceArg{
		base:  base,
		value: values,
	}
}

func (s *intSliceArg) Value() interface{} {
	return s.value
}

func (s *intSliceArg) ParseFlag() app.Flag {
	return &cli.IntSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewIntSlice(s.Value().([]int)...),
		DefaultText: s.DefaultText(),
	}
}
