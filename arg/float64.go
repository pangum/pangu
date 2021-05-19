package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*float64Arg)(nil)

type float64Arg struct {
	*base

	// 值
	value float64
}

// NewFloat64 创建一个浮点型参数
func NewFloat64(base *base, value float64) *float64Arg {
	return &float64Arg{
		base:  base,
		value: value,
	}
}

func (i *float64Arg) Value() interface{} {
	return i.Value
}

func (i *float64Arg) ParseFlag() app.Flag {
	return &cli.Float64Flag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       i.Value().(float64),
		DefaultText: i.DefaultText(),
	}
}
