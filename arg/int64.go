package arg

import (
	`github.com/pangum/pangu/app`
	`github.com/urfave/cli/v2`
)

var (
	_         = NewInt64
	_ app.Arg = (*int64Arg)(nil)
)

type int64Arg struct {
	*base

	// 值
	value int64
}

// NewInt64 创建一个整形参数
func NewInt64(base *base, value int64) *int64Arg {
	return &int64Arg{
		base:  base,
		value: value,
	}
}

func (i *int64Arg) Value() interface{} {
	return i.Value
}

func (i *int64Arg) Flag() app.Flag {
	return &cli.Int64Flag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       i.Value().(int64),
		DefaultText: i.DefaultText(),
	}
}
