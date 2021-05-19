package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*int64SliceArg)(nil)

type int64SliceArg struct {
	arg

	// 值
	value []int64
}

// NewInt64Slice 创建一个整形数组参数
func NewInt64Slice(arg arg, values ...int64) *int64SliceArg {
	return &int64SliceArg{
		arg:   arg,
		value: values,
	}
}

func (s *int64SliceArg) Value() interface{} {
	return s.value
}

func (s *int64SliceArg) ParseFlag() app.Flag {
	return &cli.Int64SliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewInt64Slice(s.Value().([]int64)...),
		DefaultText: s.DefaultText(),
	}
}
