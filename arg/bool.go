package arg

import (
	`github.com/storezhang/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*boolArg)(nil)

type boolArg struct {
	*base

	// 值
	value bool
}

// NewBool 创建布尔参数
func NewBool(base *base, value bool) *boolArg {
	return &boolArg{
		base:  base,
		value: value,
	}
}

func (b *boolArg) Value() interface{} {
	return b.Value
}

func (b *boolArg) ParseFlag() app.Flag {
	return &cli.BoolFlag{
		Name:        b.Name(),
		Aliases:     b.Aliases(),
		Usage:       b.Usage(),
		Value:       b.Value().(bool),
		DefaultText: b.DefaultText(),
	}
}
