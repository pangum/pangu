package arg

import (
	`github.com/pangum/pangu/app`
	`github.com/urfave/cli/v2`
)

var (
	_         = NewInt64s
	_ app.Arg = (*int64sArg)(nil)
)

type int64sArg struct {
	*base

	// 值
	values []int64
}

// NewInt64s 创建一个整形数组参数
func NewInt64s(base *base, values ...int64) *int64sArg {
	return &int64sArg{
		base:   base,
		values: values,
	}
}

func (i *int64sArg) Value() interface{} {
	return i.values
}

func (i *int64sArg) Flag() app.Flag {
	return &cli.Int64SliceFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       cli.NewInt64Slice(i.Value().([]int64)...),
		DefaultText: i.DefaultText(),
	}
}
