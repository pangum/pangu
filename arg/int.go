package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*intArg)(nil)

type intArg struct {
	*base

	// 值
	value int
}

// NewInt 创建一个整形参数
func NewInt(base *base, value int) *intArg {
	return &intArg{
		base:  base,
		value: value,
	}
}

func (i *intArg) Value() interface{} {
	return i.value
}

func (i *intArg) ParseFlag() app.Flag {
	return &cli.IntFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       i.Value().(int),
		DefaultText: i.DefaultText(),
	}
}
