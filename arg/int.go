package arg

import (
	"github.com/urfave/cli/v2"

	"github.com/pangum/pangu/app"
)

var (
	_         = NewInt
	_ app.Arg = (*intArg)(nil)
)

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

func (i *intArg) Flag() app.Flag {
	return &cli.IntFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Value:       i.Value().(int),
		DefaultText: i.DefaultText(),
	}
}
