package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*int64Arg)(nil)

type int64Arg struct {
	arg

	// 值
	value int64
}

// NewInt64 创建一个整形参数
func NewInt64(arg arg, value int64) *int64Arg {
	return &int64Arg{
		arg:   arg,
		value: value,
	}
}

func (i *int64Arg) Value() interface{} {
	return i.Value
}

func (i *int64Arg) ParseFlag() app.Flag {
	return &cli.Int64Flag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       i.Value().(int64),
		DefaultText: i.DefaultText(),
	}
}
